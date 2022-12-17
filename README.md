# tgbot api types

A list of [Telegram Bot API types](https://core.telegram.org/bots/api)

### Installation 

```

go get github.com/Laaguini/tgbot-types

```

### Usage example

```go

import (
    types "github.com/Laaguini/tgbot-types"
    "net/http"
    "json"
    "fmt"
)

func main() {
    var UpdateResult types.UpdateResult

    res, _ := http.get("https://api.telegram.org/bot<YOUR_TOKEN>/getUpdates")

    defer res.Body.Close()

    body, _ := io.ReadAll(res.Body)
    json.Unmarshal(body, &UpdateResult)

    if UpdateResult.Ok {
        fmt.Println(UpdateResult.Updates.Length)
    }
}

```