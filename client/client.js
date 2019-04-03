const {GetDiaryRequest} = require("./proto/english_diary_pb.js");
const {EnglishDiaryClient} = require("./proto/english_diary_grpc_web_pb.js");

var client = new EnglishDiaryClient("http://localhost:9090");
var request = GetDiaryRequest();

client.getDiary(undefined, undefined, (err, response) => {
    console.log(response);
})