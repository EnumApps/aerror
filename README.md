# aerror
Traceable Error for Golang

#usage 

#on error
	if err != nil {
    return aerror.WrapError(err)
  }
		
#on log/output
  fmt.Println(err) // use as if normal error

  switch a:= err.(type){
    case *aerror.AError:
      fmt.Println(a.Trace)
      break
    default:
      fmt.Println(err)
      break
  }