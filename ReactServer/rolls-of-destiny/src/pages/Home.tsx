import "../css/Home.css";
import { useState } from "react";
import TopAppBar from "../bars/TopAppBar";
import { Link } from "react-router-dom";
import Button from "@mui/material/Button";
import HomeScreenButtonGroup from "../components/homeScreenButtonGroup";
import TextField from "@mui/material/TextField";
import * as React from "react";
import Box from "@mui/material/Box";
import { profile } from "../types/profileTypes";
import { Modal } from "@mui/material";
import testImage from "../soundtracks/testImage.png"

export default function Home({ loggedIn, setLoggedIn }: { loggedIn: boolean, setLoggedIn: Function }) {
    const [playOpened, setPlayOpened] = useState<boolean>(false);

    const users: profile[] = [
        { username: "Bernd", rating: 839, picture: testImage, biography: "Bio for Bernd" },
        { username: "Anna", rating: 902, picture: testImage, biography: "Bio for Anna" },
        { username: "Carlos", rating: 756, picture: "https://via.placeholder.com/100", biography: "Bio for Carlos" },
        { username: "Diana", rating: 820, picture: testImage, biography: "Bio for Diana" },
        { username: "Edward", rating: 890, picture: "https://via.placeholder.com/100", biography: "Bio for Edward" }
    ];

    const [searchBar, setSearchBar] = useState("");
    const [searchResults, setSearchResults] = useState<profile[]>([]);
    const [resultsFound, setResultsFound] = useState(false);
    const [noResultsFound, setNoResultsFound] = useState(false);

    const relocate = () => {
        window.location.href = "/profile";
        setLoggedIn(!loggedIn);
    };

    function visibleButtons() {
        if (playOpened) {
            return (
                <>
                    <HomeScreenButtonGroup setPlayOpened={setPlayOpened} playOpened={playOpened} />
                </>
            );
        } else {
            return (
                <Button variant="contained" color="secondary" onClick={() => setPlayOpened(!playOpened)}> Play </Button>
            );
        }
    }

    function submitSearchBar() {
        // for connecting with backend
        /* fetch("http://localhost:9090/users", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({searchText: searchBar})
        }).then(r => {
            if (r.status === 404) {
                return r.json();
            }
        }); */

        // but now with dummy data
        const results = users.filter(user => user.username.toLowerCase().includes(searchBar.toLowerCase()));
        results.sort((a, b) => a.username.localeCompare(b.username));
        setSearchResults(results);
        if (results.length > 0) {
            setResultsFound(true);
            setNoResultsFound(false);
        } else {
            setResultsFound(false);
            setNoResultsFound(true);
        }
    }

    const handleCloseSearchModul = () => {
        setSearchBar("");
        setSearchResults([]);
        setResultsFound(false);
        setNoResultsFound(false)
    };

    const addToFriendlist = (username: string) => {
        fetch("http://localhost:9090/addToFriendlist", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({username: username})
        }).then(r => {
            return r.json()
        })
    };

    return (
        <>
            <header>
                <TopAppBar loggedIn={loggedIn} />
            </header>
            <div className="homepage">
                <div className="homeText">
                    <h1>Rolls of Destiny</h1>
                    <h3>A game made by</h3>
                    <p className={"contributor"}><Link to={"https://github.com/Ign1tz"}>Moritz Pertl</Link></p>
                    <p className={"contributor"}><Link to={"https://github.com/LukasBrezina"}>Lukas Brezina</Link></p>
                    <p className={"contributor"}><Link to={"https://github.com/Sweisser7"}>Simon Weisser</Link></p>
                </div>
                <div className="homeButtons">
                    {visibleButtons()}
                </div>
                <Box id="searchBox">
                    <TextField id="filled-basic" label="Search for a player" variant="filled"
                               value={searchBar}
                               onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                                   setSearchBar(event.target.value);
                               }} />
                    <Button variant="contained" color="primary" onClick={submitSearchBar}>Search</Button>
                </Box>
                <Modal open={resultsFound || noResultsFound} onClose={handleCloseSearchModul}>
                    <Box className="searchResultsModal">
                        <div className="modalHeader">
                            <h2>Search Results</h2>
                            <Button color="error" variant="contained" onClick={handleCloseSearchModul}>Close</Button>
                        </div>
                        <div className="results">
                            { resultsFound && searchResults.map(profile => (
                                <Box key={profile.username} className="profileBoxHome">
                                    <img src={profile.picture} alt="profile picture" className="profilePictureHome" />
                                    <div className="profileDetailsHome">
                                        <h3>{profile.username}</h3>
                                        <p>Rating: {profile.rating}</p>
                                    </div>
                                    <div className={"addFriendButton"}>
                                        <Button variant="contained" color="primary" onClick={() => addToFriendlist(profile.username)}>Add to Friendlist</Button>
                                    </div>
                                </Box>
                            ))}
                            { noResultsFound && (
                                <Box>
                                    <h3>No results found.</h3>
                                </Box>
                            )}
                        </div>
                    </Box>
                </Modal>
            </div>
            <footer style={{ textAlign: "center", fontSize: "x-small" }}>
                Copyright
            </footer>
        </>
    );
}
