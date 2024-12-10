package server

import (
	"net"
)

func SendAsciiArt(conn net.Conn) {
	// Linux ASCII art
	conn.Write([]byte("         _nnnn_\n"))
	conn.Write([]byte("        dGGGGMMb\n"))
	conn.Write([]byte("       @p~qp~~qMb\n"))
	conn.Write([]byte("       M|@||@) M|\n"))
	conn.Write([]byte("       @,----.JM|\n"))
	conn.Write([]byte("      JS^\\__/  qKL\n"))
	conn.Write([]byte("     dZP        qKRb\n"))
	conn.Write([]byte("    dZP          qKKb\n"))
	conn.Write([]byte("   fZP            SMMb\n"))
	conn.Write([]byte("   HZM            MMMM\n"))
	conn.Write([]byte("   FqM            MMMM\n"))
	conn.Write([]byte(" __| \".        |\\dS\"qML\n"))
	conn.Write([]byte(" |    `.       | `' \\Zq\n"))
	conn.Write([]byte("_)      \\.___.,|     .'\n"))
	conn.Write([]byte("\\____   )MMMMMP|   .'\n"))
	conn.Write([]byte("     `-'       `--'\n"))
}
