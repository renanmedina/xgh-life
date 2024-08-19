package gohorse

import "golang.org/x/exp/maps"

const DEFAULT_LANGUAGE = "pt-BR"

var AXIOMS = map[string]map[int]Axiom{
	"pt-BR": {
		1:  {1, "Pensou, não é XGH.", "XGH não pensa, faz a primeira coisa que vem à mente. Não existe segunda opção, a única opção é a mais rápida."},
		2:  {2, "Existem 3 formas de se resolver um problema, a correta, a errada e a XGH, que é igual à errada, só que mais rápida.", "XGH é mais rápido que qualquer metodologia de desenvolvimento de software que você conhece (Vide Axioma 14)."},
		3:  {3, "Quanto mais XGH você faz, mais precisará fazer.", "Para cada problema resolvido usando XGH, mais uns 7 são criados. Mas todos eles serão resolvidos da forma XGH. XGH tende ao infinito."},
		4:  {4, "XGH é totalmente reativo.", "Os erros só existem quando aparecem."},
		5:  {5, "XGH vale tudo, só não vale dar o toba.", "Resolveu o problema? Compilou? Commit e era isso."},
		6:  {6, "Commit sempre antes de update.", "Se der merda, a sua parte estará sempre correta.. e seus colegas que se fodam."},
		7:  {7, "XGH não tem prazo.", "Os prazos passados pelo seu cliente são meros detalhes. Você SEMPRE conseguirá implementar TUDO no tempo necessário (nem que isso implique em acessar o BD por um script malaco)."},
		8:  {8, "Esteja preparado para pular fora quando o barco começar a afundar… ou coloque a culpa em alguém ou algo.", "Pra quem usa XGH, um dia o barco afunda. Quanto mais o tempo passa, mais o sistema vira um monstro. O dia que a casa cair, é melhor seu curriculum estar cadastrado na APInfo, ou ter algo pra colocar a culpa."},
		9:  {9, "Seja autêntico, XGH não respeita padrões.", "Escreva o código como você bem entender, se resolver o problema, commit e era isso."},
		10: {10, "Não existe refactoring, apenas rework.", "Se der merda, refaça um XGH rápido que solucione o problema. O dia que o rework implicar em reescrever a aplicação toda, pule fora, o barco irá afundar (Vide Axioma 8)."},
		11: {11, "XGH é totalmente anárquico.", "A figura de um gerente de projeto é totalmente descartável. Não tem dono, cada um faz o que quiser na hora que os problemas e requisitos vão surgindo (Vide Axioma 4)."},
		12: {12, "Se iluda sempre com promessas de melhorias.", "Colocar TODO no código como uma promessa de melhoria ajuda o desenvolvedor XGH a não sentir remorso ou culpa pela cagada que fez. É claro que o refactoring nunca será feito (Vide Axioma 10)."},
		13: {13, "XGH é absoluto, não se prende à coisas relativas.", "Prazo e custo são absolutos, qualidade é totalmente relativa. Jamais pense na qualidade e sim no menor tempo que a solução será implementada, aliás… não pense, faça!"},
		14: {14, "XGH é atemporal.", "Scrum, XP… tudo isso é modinha. O XGH não se prende às modinhas do momento, isso é coisa de viado. XGH sempre foi e sempre será usado por aqueles que desprezam a qualidade."},
		15: {15, "XGH nem sempre é POG.", "Muitas POG's exigem um raciocínio muito elevado, XGH não raciocina (Vide Axioma 1)."},
		16: {16, "Não tente remar contra a maré.", "Caso seus colegas de trabalho usam XGH para programar e você é um coxinha que gosta de fazer as coisas certinhas, esqueça! Pra cada Design Pattern que você usa corretamente, seus colegas gerarão 10 vezes mais código podre usando XGH."},
		17: {17, "O XGH não é perigoso até surgir um pouco de ordem.", "Este axioma é muito complexo, mas sugere que o projeto utilizando XGH está em meio ao caos. Não tente por ordem no XGH (Vide Axioma 16), é inútil e você pode jogar um tempo precioso no lixo. Isto fará com que o projeto afunde mais rápido ainda (Vide Axioma 8). Não tente gerenciar o XGH, ele é auto suficiente (Vide Axioma 11), assim como o caos."},
		18: {18, "O XGH é seu brother, mas é vingativo.", "Enquanto você quiser, o XGH sempre estará do seu lado. Mas cuidado, não o abandone. Se começar um sistema utilizando XGH e abandoná-lo para utilizar uma metodologia da moda, você estará fudido. O XGH não permite refactoring (vide axioma 10), e seu novo sistema cheio de frescurites entrará em colapso. E nessa hora, somente o XGH poderá salvá-lo."},
		19: {19, "Se tiver funcionando, não rela a mão.", "Nunca altere, e muito menos questione um código funcionando. Isso é perda de tempo, mesmo porque refactoring não existe (Vide Axioma 10). Tempo é a engrenagem que move o XGH e qualidade é um detalhe desprezível."},
		20: {20, "Teste é para os fracos.", "Se você meteu a mão num sistema XGH, é melhor saber o que está fazendo. E se você sabe o que está fazendo, vai testar pra que? Testes são desperdício de tempo, se o código compilar, é o suficiente."},
		21: {21, "Acostume-se ao sentimento de fracasso iminente.", "O fracasso e o sucesso andam sempre de mãos dadas, e no XGH não é diferente. As pessoas costumam achar que as chances do projeto fracassar utilizando XGH são sempre maiores do que ele ser bem sucedido. Mas sucesso e fracasso são uma questão de ponto de vista. O projeto foi por água abaixo mas você aprendeu algo? Então pra você foi um sucesso!"},
		22: {22, "O problema só é seu quando seu nome está no Doc da classe.", "Nunca ponha a mão numa classe cujo autor não é você. Caso um membro da equipe morra ou fique doente por muito tempo, o barco irá afundar! Nesse caso, utilize o Axioma 8."},
		23: {23, "Mais é mais.", "Com o XGH você prospera na duplicação de código - a qualidade do código não tem sentido e não há tempo para revisões de código ou refatoração. O tempo é essencial, então copie e cole, rapidamente!"},
	},
	"en-US": {
		1:  {1, "I think therefore it's not XGH.", "In XGH you don't think, you do the first thing that comes to your mind. There's not a second option as the first one is faster."},
		2:  {2, "There are 3 ways of solving a problem: the right way, the wrong way and the XGH way which is exactly like the wrong one but faster.", "XGH is faster than any development process you know (see Axiom 14)."},
		3:  {3, "You'll always need to do more and more XGH.", "For every solved problem using XGH 7 more are created. And all of them will be solved using XGH. Therefore XGH tends to the infinite."},
		4:  {4, "XGH is completely reactive.", "Errors only come to exist when they appear."},
		5:  {5, "In XGH anything goes.", "It solves the problem? It compiled? You commit and don't think about it anymore."},
		6:  {6, "You commit always before updating.", "If things go wrong your part will always be correct... and your colleagues will be the ones dealing with the problems."},
		7:  {7, "XGH don't respect timelines.", "Timelines given to you by your clients are all but important. You will ALWAYS be able to implement EVERYTHING in time (even if that means accessing the DB through some crazy script)."},
		8:  {8, "Be ready to jump off when the boat starts sinking. Or blame someone else.", "For people using XGH someday the boat sinks. As time passes by the system grows into a bigger monster. You better have your resume ready for when the thing comes down. Or have someone else to blame."},
		9:  {9, "Be authentic. XGH don't follow patterns.", "Write code as you may want. If it solves the problem you must commit and forget about it."},
		10: {10, "There's no refactoring just rework.", "If things ever go wrong just use XGH to quickly solve the problem. Whenever the problem requires rewriting the whole software it's time for you to drop off before the whole thing goes down."},
		11: {11, "XGH is anarchic.", "There's no need for a project manager. There's no owner and everyone does whatever they want when the problems and requirements appear."},
		12: {12, "Always believe in improvement promises.", "Putting TODO comments in the code as a promise that the code will be improved later helps the XGH developer. He/She won't feel guilt for the shit he/she did. Sure there won't be no refactoring (see Axiom 10)."},
		13: {13, "XGH is absolute, it doesnt attaches to relative things.", "Delivery dates and costs are absolute things. Quality is relative. Never think about quality but instead think about the minimum time required to implement a solution. Actually, don't think. Do!"},
		14: {14, "XGH is not a fad.", "Scrum, XP? Those are just trends. XGH developers don't follow temporary trends. XGH will always be used by those who despise quality."},
		15: {15, "XGH is not always WOP (Workaround-oriented programming).", "Many WOP require smart thinking. XGH requires no thinking (see Axiom 1)."},
		16: {16, "Don't try to row against the tide.", "If your colleagues use XGH and you are the only sissy who wants to do things the right way then quit it! For any design pattern that you apply correctly your colleagues will generate 10 times more rotten code using XGH."},
		17: {17, "XGH is not dangerous until you see some order in it.", "This axiom is very complex but it says that a XGH project is always in chaos. Don't try to put order into XGH (see Axiom 16). It's useless and you'll spend a lot of precious time. This will make things go down even faster. Don't try to manage XGH as it's auto-sufficient (see Axiom 11) as it's also chaos."},
		18: {18, "XGH is your bro. But it's vengeful.", "While you want it XGH will always be at your side. But be careful not to abandon him. If you start something using XGH and then turn to some trendy methodology you will be fucked up. XGH doesn't allow refactoring (see Axiom 10) and your new sissy system will collapse. When that happens only XGH can save you."},
		19: {19, "If it's working don't bother.", "Never ever change - or even think of question - a working code. That's a complete waste of time even more because refactoring doesn't exist (see Axiom 10). Time is the engine behind XGH and quality is just a meaningless detail."},
		20: {20, "Tests are for pussies.", "If you ever worked with XGH you better know what you're doing. And if you know what you're doing why test then? Tests are a waste of time. If it compiles it's good."},
		21: {21, "Be used to the 'living on the edge' feeling.", "Failure and sucess are really similar and XGH is not different. People normally think that a project can have greater chances of failing when using XGH. But success is just a way of seeing it. The project failed. You learned something with it? Then for you it was a success!"},
		22: {22, "The problem is only yours when you name is on the code docs.", "Never touch a class of code which you're not the author. When a team member dies or stays away for too long the thing will go down. When that happens use Axiom 8."},
		23: {23, "More is more.", "With XGH you thrive on code duplication - code quality is meaningless and there's no time for code reviews or refactoring.  Time is of the essence, so copy and paste, quickly!"},
	},
}

func IsLanguageSupported(lang string) bool {
	keys := maps.Keys(AXIOMS)

	for _, key := range keys {
		if key == lang {
			return true
		}
	}

	return false
}
