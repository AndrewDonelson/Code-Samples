# Auto Jigger
This project is many years old, just as HTML5 was being tested but was not officially released. There really was no other frameworks at the time. As I do not use react or vue - I do still use this for any of my projects that need basic webapp feature because it is easy, light and fast.

A simple tool for web developers that properly and automatically handles the inclusion of Modernizr, jQuery, Bootstrap and Font-Awesome into web pages providing the best possible performance with lowest bandwidth.

## Features
Simple, Small, Fast - Great for embeding as ou will not have to have Bootstrap, jQuery, Font-Awesome or Modernizr in your repository or container as they are pulled from CDN.
saving 21 characters per refernece (uncompiled)

### Javascript
Smaller JS footprint as it uses shortcuts for ofton used common referneces. For example Shorthand access to window.document.head[0] is H$ 

### CSS
extended bootstrap by adding
- Responsive helper classes for Orientation 
-- hidden-[portrait||landscape]
-- visible-[portrait||landscape]
- Mobile Device Width (.mdw) through css
- Dots Per Inch (.dpi) through css
- .autoScale will atomatically scale images based on orientation and dpi
- Simple classes for hiding/revealing elements (.vanish && .appear) - Not the same as Bootstrap.
- Simple cross platform opacity classes [opacity0, opacity20, opacity40, opacity60, opacity80, opacity100]

###HTML
Easy to use, only 3 div's (logo, wrapper and tail)
