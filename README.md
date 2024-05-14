Fast loading and easy to use drawing board application. Built with raylib bindings for GO.

DEBi is designed to be fast to open and easy to use with minimal ui. Everything is handled with hotkeys such as switching color and clearing the screen.

The main branch version of DEBi paints shapes to the screen without frambuffer. This makes the progam easy to build but limits custom cursors.
This branch is for rebuilding the system to allow for custom cursors with a real updated framebuffer.
This will be done by using raylibs TextureMode 

Known issues 
    1. Screen may flicker due to framebuffer issues I haven't figured out yet
