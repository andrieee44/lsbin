# LSBIN

[NAME](#NAME)  
[SYNOPSIS](#SYNOPSIS)  
[DESCRIPTION](#DESCRIPTION)  
[EXAMPLE](#EXAMPLE)  
[SEE ALSO](#SEE%20ALSO)  
[AUTHOR](#AUTHOR)  

------------------------------------------------------------------------

## NAME <span id="NAME"></span>

lsbin − list binaries in JSON format

## SYNOPSIS <span id="SYNOPSIS"></span>

**lsbin**

## DESCRIPTION <span id="DESCRIPTION"></span>

**lsbin** lists all binaries in **\$PATH** in JSON format. It is
originally designed to work with *cmenu*(1). The output will be
**map\[\<BINARY NAME\>\]\<FULL PATH\>**.

## EXAMPLE <span id="EXAMPLE"></span>

Create a menu for binaries like in *dmenu*(1):

**\$ lsbin \| cmenu fzf**

## SEE ALSO <span id="SEE ALSO"></span>

***cmenu***(1) *dmenu*(1) *fzf*(1)

## AUTHOR <span id="AUTHOR"></span>

Kris Andrie Ortega (andrieee44@gmail.com)

------------------------------------------------------------------------
