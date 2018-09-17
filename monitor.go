package main

import (
    "fmt"
    "time"
)

func Monitor(ctx *SessionManagerContext, nsgs []string) {
    ticker1 := time.NewTicker(30*time.Second)
    go processTicker(ctx,ticker1)
    SetNsgs(ctx,nsgs)
    WatchNsgs(ctx)
}

func processTicker(ctx *SessionManagerContext, ticker *time.Ticker) {
    for range ticker.C {
        updateRIB(ctx)
        fmt.Printf("ticker \n")
    }
}

func updateRIB(ctx *SessionManagerContext) {
    ctx.azure.GetSubnets("")
    //TODO: remove entries from rib if they are not in this list (take nexthop into consideration)
    //TODO: add entries in rib if they are not already in it (take nexthop into consideration)
}

func onPrimaryNsgChanged(nsg *nsg, ctx *SessionManagerContext) {
    if (nsg != nil) {
        fmt.Printf("------> PRIMARY NSG IS NOW %s <------\n",nsg.address)
        if (nsg.IsActiveSpeaker(ctx)) {
            ctx.azure.ChangeUplink(nsg.address)
            updateRIB(ctx)
        }
    } else {
        fmt.Printf("------> ALL NSGs ARE IN STANDBY <------\n")
    }
}