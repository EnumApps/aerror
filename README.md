# aerror
Traceable Error for Golang
[![GoDoc](https://godoc.org/github.com/EnumApps/aerror?status.svg)](http://godoc.org/github.com/EnumApps/aerror)

#usage 

#on error
      if err != nil {
        return aerror.WrapError(err)
      }
		
#on log/output
      fmt.Println(err) // use as a normal error , it will always compatable

      switch a:= err.(type){
        case *aerror.AError:
          fmt.Println(a.Trace)
          break
        default:
          fmt.Println(err)
          break
      }

#use a premade tool

      import "github.com/EnumApps/aerror/debugutil"

      if err != nil {
        debugutil.PrintTrace(err)
      }