win := createWindow()
initGeom := win.Geometry()
screenGeom := X.RootWin().Geometry()

_ = ewmh.WmStateReqExtra(
    X, win.Id, ewmh.StateAdd,
    "_NET_WM_STATE_MAXIMIZED_HORZ", "", 2,
)
waitForConfigureNotify()
newGeom = win.Geometry()
assertGeomEquals(
    xrect.New(
        screenGeom.X(), initGeom.Y(),
        screenGeom.Width(), initGeom.Height(),
    ),
    newGeom,
    "invalid geometry",
)