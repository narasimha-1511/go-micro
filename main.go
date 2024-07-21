

package main

import (
	"context"
	"fmt"
	"github.com/narasimha-1511/go-micro/application"

	"os"
	"os/signal"
)

// to start the redis server in windoes -> 
// run it on wsl -> 
// follow the redis documentation from the official website

func main(){	
	// TODO: Implement
	app:= appilication.New();
	
	ctx ,cancel := signal.NotifyContext(context.Background(),os.Interrupt);

	defer cancel();

	error:= app.Start(ctx);

	if error!=nil{
		fmt.Println("Error starting the application",error);
	}

}
