XGetWindowProperty(
    dpy, win, atom, 0, 0, False,
    AnyPropertyType, &type_ret, &format_ret,
    &num_ret, &bytes_after, &prop_ret
);