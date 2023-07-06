# goworkrbx
Stop spending time manually setting up your Rojo workflow
\
 With this pure Go CLI you can configure your Rojo/Aftman workflow **once** and let goworkrbx lift the heavy working everytime you want to make a new project!

 ## Getting Started
 To begin configuring your workflow use the following template to configure

 ```bash
 goworkrbx [setting] -e [bool]
 ```

 ## Configuring

 Here is the current list of avaiable settings:

 ```bash
 goworkrbx rojo -e [bool] ## you should enable this in 90% of use cases

 goworkrbx wally -e [bool] ## should it create a new wally.toml?

 goworkrbx aftman -e [bool] ## should it create a new aftman.toml?

 goworkrbx aftman add [name/scope] # tool to add to the aftman.toml

 goworkrbx aftman remove [name/scope] # remove tools from the default aftman.toml
 ```

 ### **Important configuration disclaimer**

 Currently (1.1.0) if you are enabling a tool, e.g. wally, you MUST add it to the aftman tools. The reason it does not automatically install that tool is to have fork support. This tool was created for my personal projects and I use the UpliftGame fork of Rojo as it has some bug fixes not yet merged to the main Rojo.

