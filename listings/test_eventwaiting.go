// waits for event satisfying the test function
func waitForEvent(test func(event xgb.Event) bool) {
    timeout := time.After(1 * time.Second)
    for {
        select {
        case <-timeout:
            return
        default:
            for {
                event, err := X.Conn().PollForEvent()
                if test(event) {
                    return
                }
                if event == nil && err == nil {
                    break
                }
            }
        }
    }
}

func waitForPropertyChange(id xproto.Window, atomName string) {
    waitForEvent(func(event xgb.Event) bool {
        e, ok := event.(xproto.PropertyNotifyEvent)
        if ok {
            atom, _ := xprop.Atm(X, atomName)
            return e.Window == id && atom == e.Atom
        }
        return false
    })
}

waitForPropertyChange(X.RootWin(), "_NET_NUMBER_OF_DESKTOPS")