package main
import(
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"

)

func main(){

apiKey:= "AIzaSyBkwlPSthUX0vbI_ltQU8z5ThnOcyKiKrE"
model:="gemini-1.5-pro"
apiURL:=fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/%s:generateContent?key=%s",model,apiKey)


reader:=bufio.NewReader(os.Stdin)
fmt.Print("Ask a Question: ")
question , err:=reader.ReadString('\n')
if err!=nil{
    log.Fatalf("Error reading input:%v",err)
}

question = strings.TrimSpace(question)



requestBody:=map[string]interface{}{
    "contents":[]map[string]interface{}{
        {"parts":[]map[string]string{
            {"text":question},
    }},
},
}

jsonData,err:=json.Marshal(requestBody)
if err!=nil{
    log.Fatalf("Error creating json files:%v",err)
}


req ,err:=http.NewRequest("POST",apiURL,bytes.NewBuffer(jsonData))

if err!=nil{
    log.Fatalf("Error creating request POST:%v",err)
}
req.Header.Set("Content-type","application/json")


client :=&http.Client{}
resp,err :=client.Do(req)
if err!=nil{
    log.Fatalf("Error sending request:%v",err)
}
defer resp.Body.Close()






body ,err:=ioutil.ReadAll(resp.Body)
if err!=nil{
    log.Fatalf("Error reading response:%v",err)
}

outputFile :="output.txt"
err=ioutil.WriteFile(outputFile,body,0644)
if err!=nil{
    log.Fatalf("Error writing to file:%v",err)
}
fmt.Printf("\n Response saved in '%s' \n",outputFile)

}

