cookie = xcb_get_property(
    dpy, False, win, atom, XCB_GET_PROPERTY_TYPE_ANY, 0, 0
);
// do something while waiting for the response
reply = xcb_get_property_reply(dpy, cookie, NULL);