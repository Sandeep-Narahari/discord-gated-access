package bot

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
)

func (m *Module) Channels() {
	fmt.Println("Session====================", m.session)
	c, err := m.session.GuildChannels("1010831009010425906")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, val := range c {
		fmt.Println(val.ID, val.Name, val.NSFW)
	}
}

func (m *Module) AddHandleMsgCreate() {

	m.session.AddHandler(func(s *discordgo.Session, m *discordgo.ChannelUpdate) {
		if m.Channel.Type == 4 {
			p, err := s.Channel(m.Channel.ID)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(p.Name, p.ID)
			for _, val := range m.Channel.PermissionOverwrites {
				// fmt.Println(val.ID)
				r, err := s.State.Role(m.GuildID, val.ID)
				if err != nil {
					panic(err.Error())
				}
				fmt.Println(r.ID, r.Name)
			}
		}
	})
}

func (m *Module) GetAddressAndValidate() {
	m.session.AddHandler(func(s *discordgo.Session, c *discordgo.MessageCreate) {
		if c.Author.ID == s.State.User.ID {
			return
		}
		channelMsg, err := s.ChannelMessage("1045318160791912530", c.Message.ID)
		if err != nil {
			panic(err.Error())
		}
		if strings.HasPrefix(channelMsg.Content, "autonomy") && len(channelMsg.Content) == 47 {
			s.ChannelMessageSendComplex("1045318160791912530", &discordgo.MessageSend{Content: "we are verifying your wallet address" + c.Author.Mention()})
			fmt.Println(c.Author.ID)
			//Database Query to check
			var address string = m.db.CheckDiscordId(c.Author.ID)
			if utf8.RuneCountInString(address) == 0 || address == " " {
				s.ChannelMessageSendComplex("1045318160791912530", &discordgo.MessageSend{Content: "sorry your not connected with app.autonomy.network" + c.Author.Mention()})
			} else {
				//if given address matches means user connected with our app
				if address == channelMsg.Content {
					fmt.Println("Your connected with app.autonomy.network", c.Author.Mention())
					//

					//request NFT's to rpc using address

					//
					var nfts []string = m.db.GetAllNftFromAddress(address)

					fmt.Println("Here are my nfts", nfts)
					if len(nfts) > 0 {
						// Setting the role ids

						// After getting the Roleid from Gated list and discordId set the role ids
						// setRoleId(address, nfts, m.db.GetGatedList())
						var gatedList = m.db.GetGatedList()
						fmt.Println(gatedList)
					} else {
						s.ChannelMessageSendComplex("1045318160791912530", &discordgo.MessageSend{Content: "Sorry you don't have nfts , please visit app.autonomy.network" + c.Author.Mention()})
					}

				} else {
					s.ChannelMessageSendComplex("1045318160791912530", &discordgo.MessageSend{Content: "Your providing the wrong wallet address, Please check and try again" + c.Author.Mention()})
				}
			}
		} else if c.Author.Bot {
			return
		} else {
			s.ChannelMessageSendComplex("1045318160791912530", &discordgo.MessageSend{Content: "This is not the valid wallet address" + c.Author.Mention()})
		}
	})
}

func (m *Module) VerifyNft(address []string, gatedList []string) {
	var nfts = Nfts()
	var Gated = m.db.GetGatedList()
	for i, nft := range nfts {
		for j, c := range Gated {
			if strings.HasPrefix(nft, Gated[j].Collection_id) {
				var discordid = m.db.GetDiscordId(nft[i].Address)
				var roleId = Gated[j].Roleid
				m.SetRoleIdToUser(discordid, roleId)

				fmt.Println(nft, c, "Hey you got your role , Congragulations")

				break
			}
		}
	}
}
func (m *Module) SetRoleIdToUser(discordId string, roleId string) {
	err := m.session.GuildMemberRoleAdd("1010831009010425906", "1022444713753714689", "1044225010157498429")
	if err == nil {
		fmt.Println("Role id created")
	}

}
