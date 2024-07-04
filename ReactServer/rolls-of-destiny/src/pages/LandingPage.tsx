import TopAppBar from "../bars/TopAppBar";

export default function LandingPage({ loggedIn }: { loggedIn: boolean }) {
    return (
        <>
            <TopAppBar loggedIn={loggedIn} />
            <div style={{
                backgroundColor: "#55465d", // dark purple background
                minHeight: "100vh",
                padding: "20px",
                display: "flex",
                flexDirection: "column",
                alignItems: "center"
            }}>
                <h1 style={{
                    color: "#ffffff",
                    fontFamily: "'Cinzel', serif",
                    fontSize: "48px",
                    margin: "20px 0"
                }}>
                    Rolls of Destiny
                </h1>
                <div className="landingPageText" style={{
                    fontFamily: "'Times New Roman', Times, serif",
                    color: "#2c3e50",
                    padding: "20px",
                    textAlign: "justify",
                    backgroundColor: "#f4e1c1",
                    borderRadius: "10px",
                    boxShadow: "0 4px 8px rgba(0, 0, 0, 0.5)",
                    maxWidth: "800px",
                    margin: "20px",
                }}>
                    <h3>
                        In the realm of ancient prophecies and mystical realms lies a game of fate and strategy: Rolls of
                        Destiny. Inspired by the revered "Knucklebones" from the legendary "Cult of the Lamb," this game
                        elevates the classic dice duel to epic proportions. With the roll of the dice, players summon their
                        luck and wits, engaging in a battle where each number can shift the tides of victory.
                    </h3>
                    <h3>
                        But fate is not merely decided by the bones alone. Enchanted cards come into play, each imbued with
                        magical powers that alter the course of the game. Some cards bless players with fortune, others
                        curse their opponents with misfortune. Each turn becomes a dance of destiny, where strategy
                        intertwines with the whims of chance.
                    </h3>
                    <h3>
                        Gather your courage, summon your inner strategist, and prepare to face your opponents in this game
                        where every roll matters, and the next card could seal your fate. Will you rise as the champion of
                        destiny or fall to the whims of fortune? The dice are cast, and the cards await in Rolls of Destiny.
                    </h3>
                    <h6>Copyright @ChatGPT</h6>
                </div>
            </div>
        </>
    )
}
