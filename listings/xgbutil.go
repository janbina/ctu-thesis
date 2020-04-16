// WM_NORMAL_HINTS
func WmNormalHintsGet(
    xu *xgbutil.XUtil, win xproto.Window,
) (*NormalHints, error)

prop, err := xprop.GetProperty(x, win, atomName)

func PropValNums(
    reply *xproto.GetPropertyReply, err error
) ([]uint, error)

func PropValWindows(
    reply *xproto.GetPropertyReply, err error
) ([]xproto.Window, error)
