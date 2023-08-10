package tcp

import (
	"fmt"
	"net"
	"strings"
)

// change_av allows to user to put an avatar next to his name using the --av flag.
// he has the choice between 14 avatars
func (c *Client) change_av(conn net.Conn, nickname string) {
	for {
		//dislaying the avatar list
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
		// cannot enter empty input neither a negative number nor a number exceeding the list size
		if c.nickname == "" || Atoi(c.nickname) <= 0 || Atoi(c.nickname) > 14 {
			conn.Write([]byte("❌ invalid avatar try again\n"))
			name_invalid := fmt.Sprintf("❌ %s entered an invalid avatar number.\n", nickname)
			logs(name_invalid)
			fmt.Print(name_invalid)
			continue // pursue the loop
		} else if c.nickname != "" {
			switch c.nickname {
			case "1":
				c.nickname = "🧔 " + nickname
			case "2":
				c.nickname = "👩 " + nickname
			case "3":
				c.nickname = "🤴️ " + nickname
			case "4":
				c.nickname = "👸️ " + nickname
			case "5":
				c.nickname = "🧝‍♀️ " + nickname
			case "6":
				c.nickname = "👩‍💻️ " + nickname
			case "7":
				c.nickname = "🐼️ " + nickname
			case "8":
				c.nickname = "🧞‍♀️ ️" + nickname
			case "9":
				c.nickname = "🕴️ " + nickname
			case "10":
				c.nickname = "🕺️ " + nickname
			case "11":
				c.nickname = "💃️ " + nickname
			case "12":
				c.nickname = "🦩️ " + nickname
			case "13":
				c.nickname = "🦚️ " + nickname
			case "14":
				c.nickname = "🐞️ " + nickname
			default: // returning the initial nickname
				c.nickname = nickname
			}
			break

		} else {
			break
		}
	}

	change_mess := fmt.Sprintf("🔁 %s has changed his avatar to %s ...\n", nickname, c.nickname)
	c.Sendmess(change_mess, clients) //notifying the chatromm members
	logs(change_mess)                // storing activity in a log file
	fmt.Print(change_mess)           //terminal logs
}

// change allows to the user to change his name using the --nick flag
func (c *Client) change(conn net.Conn, initial string) {
	for {
		conn.Write([]byte("[ENTER YOUR NAME] ➡ : "))
		username, err := c.readInput()
		if err != nil {
			name_error := fmt.Sprintf("❌ Error while receiving the username: %s\n", err)
			logs(name_error)      // storing activity in a log file
			fmt.Print(name_error) //terminal logs
			conn.Close()          // shutting the onnection
			return
		}

		c.nickname = strings.TrimSpace(username)
		if c.nickname == "" || len(c.nickname) > 20 {
			conn.Write([]byte("❌ The name is not valid or too long please try again\n"))
			name_invalid := fmt.Sprintf("❌ User %s entered an invalid name.\n", conn.LocalAddr().String())
			logs(name_invalid)      // storing activity in a log file
			fmt.Print(name_invalid) //terminal logs
			continue                //pursue the loop
		} else {
			break
		}
	}
	change_mess := fmt.Sprintf("🔁 %s has changed his username to %s ...\n", initial, c.nickname)
	c.Sendmess(change_mess, clients) //notifying the chatromm members
	logs(change_mess)                // storing activity in a log file
	fmt.Print(change_mess)           //terminal logs
}

// remove_avatar removes the added avatar to username using the --rmav flag
func (c *Client) remove_avatar(initial string) {
	change_mess := fmt.Sprintf("%s has removed his avatar...\n", initial)
	c.nickname = initial
	c.Sendmess(change_mess, clients) //notifying the chatromm members
	logs(change_mess)                // storing activity in a log file
	fmt.Print(change_mess)           //terminal logs
}
