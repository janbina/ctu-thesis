for {
    xev, err := x.WaitForEvent()
    if err != nil {
        // handle error
        continue
    }
    switch e := xev.(type) {
    case xproto.DestroyNotifyEvent:
        // handle destroy notify event
    case xproto.PropertyNotifyEvent:
        // handle property notify event
    case xproto.ConfigureRequestEvent:
        // handle configure request event
    }
}