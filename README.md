<h3 align="center">
  Vegeta Web UI
</h3>
<h3 align="center">
  Http load testing tool
</h3>
<p align="center">
  <img src="https://raw.githubusercontent.com/yittoo/vegeta-web-ui/master/client/src/assets/gopher.png"></img>
</p>

<h3>Installation Options:</h3>

- <a href="#install_option_1">Direct bundled installation (Recommended for non-techy users)</a>
- <a href="#install_option_2">With Docker (recommended for people with experience in Command Line Interface)</a>
- <a href="#install_option_3">By bundling directly (Requires Node 12.x, Yarn ^1.19, Go ^1.14)</a>

<h4 id="install_option_1">Direct bundled installation:</h4> 

Warning this might cause trouble on AntiViruses as the software is not signed. So if possible it's best to use the Docker alternative if this occurs

- Download latest version of bundled zip from here <a href="https://github.com/yittoo/vegeta-web-ui/tree/master/dist">https://github.com/yittoo/vegeta-web-ui/tree/master/dist</a> according to operating system you use
- Unzip anywhere you'd like.
- Run the 'vegeta-web-ui' or 'vegeta-web-ui.exe' depending on which operating system you are using
- It will automatically prompt open browser http://localhost:8000 and serve the application there

<h4 id="install_option_2">With Docker:</h4> 

- (If first time using Docker) Follow the <a href="https://docs.docker.com/get-docker/" target="_blank" rel="noopener noreferrer">official Docker installation</a> for your operating system (Windows 10 Home edition is not supported) 
- Inside project directory in your terminal type the following:
<code>docker-compose up</code>
- On first installation it will install all the software necessary to build the project, it might take some time.
- Visit http://localhost:8000 in your browser

<h4 id="install_option_3">By bundling directly:</h4>

- Install Node 12.x and Yarn ^1.19 from
- Install Go ^1.14 and do default configuration
- On project directory run:
<code>
go mod download

go run . buildClient</code>
- It will build and automatically prompt open browser http://localhost:8000 and serve the application there

<h3>Contribution:</h3>

<p>Right now I'm not accepting outside help. If you have suggestions please open an issue I'll look into that as soon as possible. However feel free to fork the project and tweak as you wish, for development you can use the development docker-compose files available in the root path of the project.</p>

##### This project is built upon the original <a href="https://github.com/tsenart/vegeta" target="_blank" rel="noopener noreferrer">work of Tsenart</a>
###### Artwork by <a href="https://github.com/egonelbre/gophers" target="_blank" rel="noopener noreferrer">@egonelbre</a>
