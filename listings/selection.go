// Get selection owner of the screen we want to manage
reply := xproto.GetSelectionOwner(X.Conn(), "WM_S2").Reply()

if reply.Owner != xproto.WindowNone {
    // another client owns this selection
    // if user did not explicitly allow us to replace it,
    // we should exit here
}

// if there is no selection owner or we want to replace it,
// send the set selection owner request for our window (X.Dummy())
xproto.SetSelectionOwner(X.Conn(), X.Dummy(), "WM_S2")

// if another client owned the selection, wait for its termination
if reply.Owner != xproto.WindowNone {
    for { // timeout mechanism is omitted
        e := X.Conn().PollForEvent()
        // check that type of e is "DestroyNotifyEvent"
        if destroyEvent, ok := e.(DestroyNotifyEvent); ok {
            if destroyEvent.Window == reply.Owner {
                break
            }
        }
    }
}

// now we can register for SubstructureRedirect events
// on the root window
X.RootWin().Listen(EventMaskSubstructureRedirect);

// when another client sends the set selection owner request,
// we will receive the SelectionClearEvent, and by the ICCCM,
// we have to release managed resources and destroy the window
// that owned the selection
xevent.SelectionClearFun(func (X *XUtil, e SelectionClearEvent)) {
    if e.Selection == "WM_S2" {
        xevent.Quit(X)
    }
}).Connect(X, X.Dummy())