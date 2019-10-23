package engine

//import (
//	"encoding/json"
//	"github.com/gingerxman/eel/config"
//	"github.com/gingerxman/eel/log"
//	ali_mns "github.com/gingerxman/eel/aliyun/mns"
//)
//
//type mnsConf struct{
//	endpoint string
//	accessId string
//	accessKey string
//	topic string
//}
//
//var conf *mnsConf
//
//type mnsEngine struct{
//	engineType string
//}
//
//func newMnsEngine() *mnsEngine{
//	eg := new(mnsEngine)
//	eg.engineType = "mns"
//	return eg
//}
//
//func (this *mnsEngine) getMnsClient() ali_mns.MNSClient{
//	return ali_mns.NewAliMNSClient(conf.endpoint, conf.accessId, conf.accessKey)
//}
//
//func (this *mnsEngine) getFormattedMessage(data map[string]interface{}, tag string) ali_mns.MessagePublishRequest{
//	dataStr, _ := json.Marshal(data)
//	return ali_mns.MessagePublishRequest{
//		MessageBody: string(dataStr),
//		MessageTag: tag,
//		MessageAttributes: nil,
//	}
//}
//
//func (this *mnsEngine) Send(data map[string]interface{}, tag string){
//	client := this.getMnsClient()
//	topicName := conf.topic
//	topic := ali_mns.NewMNSTopic(topicName, client)
//
//	_, err := topic.PublishMessage(this.getFormattedMessage(data, tag))
//	if err != nil{
//		log.Logger.Error(err)
//	}
//
//}
//
//
//func init(){
//
//	eg := newMnsEngine()
//	registerEngine(eg.engineType, eg)
//
//	conf = new(mnsConf)
//	conf.accessId = config.ServiceConfig.String("aliyun::MNS_ACCESS_ID")
//	conf.accessKey = config.ServiceConfig.String("aliyun::MNS_ACCESS_KEY")
//	conf.endpoint = config.ServiceConfig.String("aliyun::MNS_ENDPOINT")
//	conf.topic = config.ServiceConfig.String("aliyun::MNS_TOPIC")
//}