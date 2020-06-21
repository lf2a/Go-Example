package main
/*
	Galois / Counter Mode é um modo de operação para cifras de blocos criptográficos de
	chave simétrica amplamente adotadas para seu desempenho. As taxas de taxa de transferência
	do GCM para canais de comunicação de alta velocidade e de ponta podem ser alcançadas com
	recursos de hardware baratos.

	um nonce é um número arbitrário que só pode ser usado uma vez. É basicamente uma palavra
	de uso único, daí o nome (N = Number (Número) e Once = Uma vez, em inglês), embora N seja
	de número, o processo também pode usar letras. Muitas vezes, é um número aleatório ou
	pseudoaleatório emitido em um protocolo de autenticação para garantir que as comunicações
	antigas não possam ser reutilizados em ataques de repetição. Eles podem também ser úteis
	como vectores de inicialização e em função hash criptográfica.
*/
