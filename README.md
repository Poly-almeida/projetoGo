# projetoGo
----------------------------------------------------------------------------------------------------------------------
Como estudo e aprendizado, este projeto, tem como base o desenvolvimento de um jogo multiplayer, onde se possa aplicar e verificar os conceitos do paralelismos computacional entre as linguagens de programação.

----------------------------------------------------------------------------------------------------------------------

Antes de mais nada, definiremos de maneira simplificada o que paralelismo na computação?

Confomre a Wikipedia acesso, 22/11/2020 <https://pt.wikipedia.org/wiki/Computa%C3%A7%C3%A3o_paralela>

Sic: "Computação paralela é uma forma de computação em que vários cálculos são realizados ao mesmo tempo,operando sob o princípio de que grandes problemas geralmente podem ser divididos em problemas menores, que então são resolvidos concorrentemente (em paralelo). Existem diferentes formas de computação paralela: em bit, instrução, de dado ou de tarefa. A técnica de paralelismo já é empregada há vários anos, principalmente na computação de alto desempenho, mas recentemente o interesse no tema cresceu devido às limitações físicas que previnem o aumento de frequência de processamento. Com o aumento da preocupação do consumo de energia dos computadores, a computação paralela se tornou o paradigma dominante nas arquiteturas de computadores sob forma de processadores multinúcleo.".

-----------------------------------------------------------------------------------------------------------------------
Conforme intruções para o desenvolvimento do projeto, se faz necessário escolher uma linguagem que não é ensinada na UNB, que possa ser comparada com outra linguagem, no quesitos de avaliação de linguagem de programação, que são: 
Legibilidade;
Capacidade de Escrita;
Confiabilidade; 
Custo;
Portabilidade; 
Paralelismo.

Neste caso o foco do projeto como informado acima, é a questão do paralelismo. Bom como a decisão de estudo foi a utilização da linguagem Goalang como ponto de focal.

Bom Goalang ou GO, como é se costuma se referir, foi desenvolvida pela empresa Google, que tem como característica ser compilada e focada em produtividade além de ser usada em *programação concorrente.

* Programa corrente é um paradigma de programação para a construção de programas que fazem o uso de execução de várias tarefas separadas ou em conjunto

Infelizmente não se pode resumir Goalang em 2 frases, mas não é o escopo do projeto, mas pode se citar sua versatilidade e capacidade para o desenvolvimento de serviço e aplicações na web, detalhe a linguagem Goalang, foi criada já com diversas bibliotecas, pacotes e protocolos que evitam que o desenvolvedor tenha que criar uma API ou usar um framework para seu trabalho. Não posso deixar de falar que um dos seus projetistas foi um dos criadores da linguagem C.

-----------------------------------------------------------------------------------------------------------------------
Depois de contextualização do projeto e linguagem em que será desenvolvida o projeto, vamos começar pelos passos iniciais:

Onde a pessoa pode achar essa linhagem, compilador, documentação e etc, simples acesse os sites abaixo:

Site da lingangem: https://golang.org/
Site para fazer download: https://golang.org/dl/
Site da documentação: https://golang.org/doc/

------------------------------------------------------------------------------------------------------------------------
 Depois de instalat a linguagem, abra o cmd (windows), terminal (linux e Mac) e digite "go".
 
 Deverá aparecer algo como abaixo, se sim instalação OK, se não, será necessário verificar as variaveis do ambiente se o "PATH" do Goalang foi criado e se está apontado para pasta certa.
 
 C:\Users\hanni>go
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

 
-------------------------------------------------------------------------------------------------------------------------------
Para trabalhar Goalang pode se usar um editor de texto simples até uma IDEA ( Integrated Development Environment ou Ambiente de Desenvolvimento Integrado), em ambos casos tenha certeza que o arquivo de GO escrito foi salvol como exemplo: nomedoarquivo.go, sim, os arquivos de Goalang tenha a terminação em com .go.

Sobre como trabalhar para ajudar na produtividade, recomendo as seguintes ferramentas de edição:

Editor de texto: Nome: Sublime TEX : site https://www.sublimetext.com/ na propria pagina já o botão de download, o site ele reconhece seu Sistema Operacional.
Após instalar, 
    pode se usar a seguinte combinação de teclas Win/Linux: ctrl+shift+p, Mac: cmd+shift+p e depois escreva Install Package Control e pressione Enter
    depois escreva GO selecione os pacotes e clique em enter.
Uma outra maneira seria r pelo Menu: 
    Click the Preferences > Browse Packages… menu
    Browse up a folder and then into the Installed Packages/ folder
    Download Package Control.sublime-package and copy it into the Installed Packages/ directory
    Restart Sublime Text

Editor de Codigo: Visual Studio Code site https://code.visualstudio.com/ na propria pagina já o botão de download, e só escolher seu Sistema Operacionl e fazer o download.

Após instalado clica nano menu esquerdo "Extensions" e digite na caixa de pesquisa "Go", selecione a extenção e instale
Uma outra forma de instalar Extensions e ir no site https://marketplace.visualstudio.com/VSCode e escolher a extenção e fazer a download e depois instalar no VSCode.

Por fim uma IDE que é paga, que ajuda na elaboração do codigo é a da jetbrains, a GoLand, site https://www.jetbrains.com/pt-br/go/

------------------------------------------------------------------------------------------------------------------------------------------------------------------

Para testar os codigos do servidor e cliente, 1º roda o codigo do servidorV2.go [ go run servidorV2go ] e depois roda em outra instancia do compilador (abre um CMD para o servidor e outro para o cliente), para roda o codigo do cliente [ go run clienteV1.go ].

    OBS: 1º compilar ou rodar o codigo do servidor e depois do cliente

---------------------------------------
Entendo o codigo:

A colocar as fenições do codigo do cliente e servidor

------------------------
* Informações pesquisadas no site <https://blog.betrybe.com/tecnologia/o-que-e-servidor/> acessado no dia 30/11.

Quais as vantagens e desvantagens de usar um servidor?

*Vantagens de usar um servidor

    ->os custos de licenciamento de software são reduzidos, já que não é necessário adquirir um para cada computador. Para grandes empresas, que operam com mais de 100    máquinas, pode ser um custo considerável;
    ->as despesas com infraestrutura também podem ser reduzidas, já que algumas aplicações podem migrar para o servidor, seja ele físico, seja em nuvem, em vez de rodar nas máquinas, evitando computadores com recursos mais robustos;
    ->aumento da segurança interna, já que os protocolos serão utilizados no servidor, e não nas pessoas usuárias. Assim, é possível centralizar as ações de segurança de dados. Além disso, é mais fácil implementar controles de acesso nos servidores, outra medida fundamental de proteção;
    ->portabilidade: a migração dos conteúdos, bem como a realização de backups, se torna mais prática, pois todas as informações estão centralizadas em um mesmo lugar;
    ->outsourcing de TI facilitado: com isso, é possível conseguir uma terceirização de serviço de fato eficiente, permitindo uma maior eficiência nas atividades realizadas por uma empresa externa.
    
*Desvantagens de usar um servidor

    ->custo desnecessário para empresas pequenas: para aquelas que operam com poucas pessoas usuárias, a aquisição de um servidor pode ser supérflua;
    ->se há uma dependência das aplicações no uso interno da empresa, caso ele passe por manutenção, a solução parará de rodar e seu negócio não poderá utilizá-lo, paralisando as atividades internas;
    ->uma falha de segurança que afete os servidores pode afetar todas as máquinas da empresa.
    
---------------------------------------
