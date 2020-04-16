cookie := xproto.GetProperty(
    conn, false, win, atom, xproto.GetPropertyTypeAny, 0, 0,
)
// do something while waiting for the response
reply, err := cookie.Reply()