in 10string,

at the time of importing multiple package, do not use curli breces.

:-  import {"fmt" 
            "strings"
            } ❌

which will leads to missing path error

instead of use

:- import ("fmt" 
            "strings"
            ) ✅

------------------------------------------------
