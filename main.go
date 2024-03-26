package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func main() {
    ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
    defer cancel()

    var nodes []*cdp.Node
    err := chromedp.Run(ctx,
        chromedp.Navigate(`https://michelinhr.wd3.myworkdayjobs.com/fr-FR/Michelin?primaryLocation=70b38acca79b016ee57eaf170819831f&primaryLocation=70b38acca79b01c63582c8170819ab1f&primaryLocation=758f4bb537a81000b449d0951a040000&jobPostingEmployeeContractType=cae294f2796f01131be50eaaa917b021&jobFamilyGroup=cae294f2796f012797e45b33f7173967&Location_Country=54c5b6971ffb4bf0b116fe7651ec789a`),
        // chromedp.WaitVisible(`.css-19uc56f`, chromedp.ByQueryAll),
        // chromedp.Nodes(`.css-19uc56f`, &nodes, chromedp.ByQueryAll),
        chromedp.WaitVisible(`.css-1q2dra3`, chromedp.ByQueryAll),
        chromedp.Nodes(`.css-1q2dra3`, &nodes, chromedp.ByQueryAll),
    )
    if err != nil {
        log.Fatal(err)
    }

    type Job struct {
        Titre string
        Date string
        Reference string
      }

    for _, node := range nodes {
        var text string
        err := chromedp.Run(ctx, chromedp.Text(node.FullXPath(), &text))
        if err != nil {
            log.Fatal(err)
        }
        // fmt.Println(strings.Split(strings.TrimSpace(text), "\n")[0])
        // fmt.Println(strings.Split(strings.TrimSpace(text), "\n")[4])
        var job Job
        //create a json object
        jobJson := `{"Titre": "` + strings.Split(strings.TrimSpace(text), "\n")[0] + `", 
        "Date": "` + strings.Split(strings.TrimSpace(text), "\n")[4] + `",
        "Reference": "` + strings.Split(strings.TrimSpace(text), "\n")[5] + `"}`
        json.Unmarshal([]byte(jobJson), &job)
        fmt.Printf("Poste: %s, Date: %s, Reference: %s", job.Titre, job.Date, job.Reference)
        fmt.Println()

    }
}