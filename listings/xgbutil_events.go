xevent.DestroyNotifyFun(
    func(x *xgbutil.XUtil, e xevent.DestroyNotifyEvent) {
        // handle destroy notify event on window *win*
    },
).Connect(X, win)

xevent.PropertyNotifyFun(
    func(x *xgbutil.XUtil, e xevent.PropertyNotifyEvent) {
        // handle property notify event on window *win*
    },
).Connect(X, win)

xevent.ConfigureRequestFun(
    func(x *xgbutil.XUtil, e xevent.ConfigureRequestEvent) {
        // handle configure request event on window *win*
    },
).Connect(X, win)