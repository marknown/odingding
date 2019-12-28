# How to use?

```
package main

import (
    "github.com/marknown/odingding"
)

func main() {
    dingding := &odingding.Dingding{
        Token : "fill your token ",
        Secret : "filll your secret",
    }

    dingding.NotifyText("you message")
    dingding.NotifyLink("your title", "your contnet", "your link", "you logo link")
}
```
