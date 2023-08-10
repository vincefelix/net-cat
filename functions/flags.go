package tcp

import (
	"fmt"
	"net"
	"strings"
)

func (c *Client) change_av(conn net.Conn, initial string) {
	for {
		conn.Write([]byte("💠 Pick an avatar:\n\n1- 🧔\n\n2- 👩‍🦰️\n\n3- 🤴️\n\n4- 👸️\n\n5- 🧝‍♀️️\n\n6- 👩‍💻️\n\n7- 🐼️ \n\n8- 🧞‍♀️️\n\n9- 🕴️\n\n10-🕺️\n\n11- 💃️\n\n12- 🦩️\n\n13- 🦚️\n\n14- 🐞️\n"))
		avatar, err := c.readInput()
		if err != nil {
			name_error := fmt.Sprintf("❌ Error while receiving the username: %s\n", err)
			logs(name_error)
			fmt.Print(name_error)
			conn.Close()
			return
		}

		c.nickname = strings.TrimSpace(avatar)
		if c.nickname == "" || Atoi(c.nickname) <= 0 || Atoi(c.nickname) > 14 {
			conn.Write([]byte("❌ invalid avatar try again\n"))
			name_invalid := fmt.Sprintf("❌ %s entered an invalid avatar number.\n", initial)
			logs(name_invalid)
			fmt.Print(name_invalid)
			continue
		} else if c.nickname != "" {
			switch c.nickname {
			case "1":
				c.nickname = "🧔 " + initial
			case "2":
				c.nickname = "👩 " + initial
			case "3":
				c.nickname = "🤴️ " + initial
			case "4":
				c.nickname = "👸️ " + initial
			case "5":
				c.nickname = "🧝‍♀️ " + initial
			case "6":
				c.nickname = "👩‍💻️ " + initial
			case "7":
				c.nickname = "🐼️ " + initial
			case "8":
				c.nickname = "🧞‍♀️ ️" + initial
			case "9":
				c.nickname = "🕴️ " + initial
			case "10":
				c.nickname = "🕺️ " + initial
			case "11":
				c.nickname = "💃️ " + initial
			case "12":
				c.nickname = "🦩️ " + initial
			case "13":
				c.nickname = "🦚️ " + initial
			case "14":
				c.nickname = "🐞️ " + initial
			}
			break

		} else {
			break
		}
	}
	change_mess := fmt.Sprintf("🔁 %s has changed his avatar to %s ...\n", initial, c.nickname)
	c.Sendmess(change_mess, clients)
	logs(change_mess)
	fmt.Print(change_mess)
}

func (c *Client) change(conn net.Conn, initial string) {
	for {
		conn.Write([]byte("[ENTER YOUR NAME] ➡ : "))
		username, err := c.readInput()
		if err != nil {
			name_error := fmt.Sprintf("❌ Error while receiving the username: %s\n", err)
			logs(name_error)
			fmt.Print(name_error)
			conn.Close()
			return
		}

		c.nickname = strings.TrimSpace(username)
		if c.nickname == "" {
			conn.Write([]byte("❌ The name is not valid please try again\n"))
			name_invalid := fmt.Sprintf("❌ User %s entered an invalid name.\n", conn.LocalAddr().String())
			logs(name_invalid)
			fmt.Print(name_invalid)
			continue
		} else {
			break
		}
	}
	change_mess := fmt.Sprintf("🔁 %s has changed his username to %s ...\n", initial, c.nickname)
	c.Sendmess(change_mess, clients)
	logs(change_mess)
	fmt.Print(change_mess)
}

func (c *Client) remove_avatar(initial string) {
	change_mess := fmt.Sprintf("%s has removed his avatar...\n", initial)
	c.nickname = initial
	c.Sendmess(change_mess, clients)
	logs(change_mess)
	fmt.Print(change_mess)
}
