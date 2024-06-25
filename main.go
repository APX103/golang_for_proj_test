package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CobraParam struct {
	Key      string `json:"key" bson:"key"`
	Type     string `json:"type" bson:"type"`
	Default  string `json:"default,omitempty" bson:"default,omitempty"`
	Required bool   `json:"required" bson:"required"`
	Help     string `json:"help,omitempty" bson:"help,omitempty"`
}

type CobraCMD struct {
	CMD    string       `json:"cmd" bson:"cmd"`
	SubCMD []CobraCMD   `json:"sub-cmd,omitempty" bson:"cmd,omitempty"`
	Params []CobraParam `json:"params,omitempty" bson:"params,omitempty"`
	Short  string       `json:"short,omitempty" bson:"short,omitempty"`
	Long   string       `json:"long,omitempty" bson:"long,omitempty"`
}

type MongoClientImpl struct {
	url    string
	db     string
	Client mongo.Client
}

func (c *MongoClientImpl) GetMongoConnectionPool() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(c.url).
		SetMaxPoolSize(uint64(10))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	c.Client = *client
}

func (c *MongoClientImpl) Find(collection string, key string, value interface{}) (*[][]byte, bool) {
	coll := c.Client.Database(c.db).Collection(collection)
	var result []bson.M
	q := bson.D{{key, value}}
	if key == "" {
		logrus.Debug("find all")
		logrus.Debugf("collection: %s", collection)
		q = bson.D{}
	}
	cursor, err := coll.Find(context.TODO(), q)
	if err != nil {
		logrus.Infof("failed to get key: %s", err)
		return nil, false
	}

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		logrus.Errorf("Decode error: %s", err)
		return nil, false
	}

	var s [][]byte
	for _, r := range result {
		jsonData, err := json.MarshalIndent(r, "", "    ")
		if err != nil {
			logrus.Errorf("failed to marshal to json []byte: %s", err)
			return nil, false
		}
		s = append(s, jsonData)
	}

	return &s, true
}

// +++++ Define type enumerate for task

type ParamEnum string

const String ParamEnum = "String"
const StringToString ParamEnum = "StringToString"

type ParamStruct struct {
	Type  ParamEnum
	Value any
}

type TaskCmd struct {
	SubCmd map[string]*TaskCmd
	Enable bool
	Params map[string]*ParamStruct
}

// +++++ End define

func NewCommand(rootCmd *cobra.Command, item *CobraCMD, taskCmd *TaskCmd) {
	_taskCmd := &TaskCmd{
		Enable: false,
		SubCmd: make(map[string]*TaskCmd),
		Params: make(map[string]*ParamStruct),
	}
	taskCmd.SubCmd[item.CMD] = _taskCmd

	_cmd := &cobra.Command{
		Use:   item.CMD,
		Short: item.Short,
		Long:  item.Long,
		Run: func(cmd *cobra.Command, args []string) {
			_taskCmd.Enable = true
		},
	}

	if len(item.Params) != 0 {
		for _, param := range item.Params {
			switch param.Type {
			case "String":
				_taskCmd.Params[param.Key] = &ParamStruct{
					Type:  ParamEnum(param.Type),
					Value: _cmd.PersistentFlags().String(param.Key, param.Default, param.Help),
				}
			case "StringToString":
				_taskCmd.Params[param.Key] = &ParamStruct{
					Type:  ParamEnum(param.Type),
					Value: _cmd.PersistentFlags().StringToString(param.Key, nil, param.Help),
				}
			default:
				fmt.Println("Type not support yet")
			}
			if param.Required {
				_cmd.MarkFlagRequired(param.Key)
			}
		}
	}

	if len(item.SubCMD) != 0 {
		for _, subCMD := range item.SubCMD {
			NewCommand(_cmd, &subCMD, _taskCmd)
		}
	}

	rootCmd.AddCommand(_cmd)
}

func PrintCommand(taskCmd *TaskCmd) {
	jenkins_cmd, ok := taskCmd.SubCmd["jenkins"]
	if !ok {
		return
	}
	fmt.Println("=========== PrintCommand ===========")
	for k, v := range jenkins_cmd.SubCmd {
		if v.Enable {
			fmt.Println("sub-cmd: " + k)
			for _, p := range v.Params {
				if p.Type == "String" {
					fmt.Println(*p.Value.(*string))
				}
				if p.Type == "StringToString" {
					fmt.Println(*p.Value.(*map[string]string))
				}
			}
		}
	}
	fmt.Println("====================================")
}

// TODO 在这里执行操作
func ExecCommand() {

}

func main() {
	c := &MongoClientImpl{}
	c.url = "mongodb://127.0.0.1:27017/super-mid"
	c.db = "super-mid"
	c.GetMongoConnectionPool()
	data, ok := c.Find("cmd_template", "cmd", "jenkins")

	if !ok {
		fmt.Println("zao!")
		panic("oh ho.")
	}

	var cmdList []*CobraCMD
	for _, d := range *data {
		group := &CobraCMD{}
		err := json.Unmarshal(d, group)
		if err != nil {
			fmt.Println("unmarshal history error")
			panic("oh ho.")
		}
		cmdList = append(cmdList, group)
	}
	// // return &funcList
	// for _, item := range cmdList {
	// 	// fmt.Println(item)
	// 	fmt.Println(item.CMD)
	// 	fmt.Println(item.Short)
	// 	fmt.Println(item.Long)
	// 	fmt.Println(item.SubCMD)
	// }

	rootCmd := &cobra.Command{
		Use:   "@Mr.meeseeks",
		Short: "Feishu Agent Build By QA Team.",
		Long:  `Feishu Agent Build By QA Team. Any question please try '--help' or ask @李佳伦 for help.`,
	}
	res := &TaskCmd{
		Enable: true,
		SubCmd: make(map[string]*TaskCmd),
		Params: make(map[string]*ParamStruct),
	}
	for _, cmd := range cmdList {
		NewCommand(rootCmd, cmd, res)
	}

	// fmt.Println("====================================")
	// rootCmd.SetArgs([]string{"--help"})
	// rootCmd.Execute()
	// fmt.Println("====================================")
	// rootCmd.SetArgs([]string{"jenkins", "--help"})
	// rootCmd.Execute()
	// fmt.Println("====================================")
	// rootCmd.SetArgs([]string{"jenkins", "build", "--help"})
	// rootCmd.Execute()
	// fmt.Println("====================================")
	rootCmd.SetArgs([]string{"jenkins", "build", "--job_name=AK47", "--param='A=k,a=K'"})
	rootCmd.Execute()
	// fmt.Println(res.SubCmd["jenkins"].SubCmd["build"])
	// fmt.Println(*res.SubCmd["jenkins"].SubCmd["build"].Params["job_name"].Value.(*string))
	// fmt.Println(*res.SubCmd["jenkins"].SubCmd["build"].Params["param"].Value.(*map[string]string))
	PrintCommand(res)
}
