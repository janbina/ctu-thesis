package window

import (
    "log"
    "time"

    "github.com/BurntSushi/xgb/xproto"
    "github.com/BurntSushi/xgbutil"
    "github.com/janbina/swm/internal/config"
    "github.com/janbina/swm/internal/decoration"
)

type Window struct {
    win         *xwindow.Window
    infoTimer   *time.Timer
    decorations decoration.Decorations
    moveState   *MoveState
}