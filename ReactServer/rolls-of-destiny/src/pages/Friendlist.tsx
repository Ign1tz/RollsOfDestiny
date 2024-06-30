import React, {useState} from 'react';
import TopAppBar from "../bars/TopAppBar";
import { profile } from "../types/profileTypes";
import '../css/Friendlist.css';
import onlineImage from "../images/green.jpeg"
import offlineImage from "../images/red.jpeg"
import testImage from "../soundtracks/testImage.png";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import {Modal} from "@mui/material";

export default function Friendlist({ loggedIn }: { loggedIn: boolean }) {

    const [searchBar, setSearchBar] = useState("");
    const [searchResults, setSearchResults] = useState<profile[]>([]);
    const [resultsFound, setResultsFound] = useState(false);
    const [noResultsFound, setNoResultsFound] = useState(false);

    const users: profile[] = [
        { username: "Bernd", rating: 839, profilePicture: testImage},
        { username: "Anna", rating: 902, profilePicture: testImage},
        { username: "Carlos", rating: 756, profilePicture: "https://via.placeholder.com/100"},
        { username: "Diana", rating: 820, profilePicture: testImage},
        { username: "Edward", rating: 890, profilePicture: "https://via.placeholder.com/100"}
    ];

    const online = true;

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
            if (r.status === 200) {
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

    const removeFromFriendlist = (username: string) => {
        fetch("http://localhost:909/removeFromFriendList", {
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
            <TopAppBar loggedIn={loggedIn} />
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
                                <img src={profile.profilePicture} alt="profile picture" className="profilePictureHome" />
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
            <div className="friendlist-container">
                <h1 className="friendlist-title">Your Friends</h1>
                <ul className="friendlist">
                    {users.map((user, index) => (
                        <li key={index} className="friendlist-item">
                            <div className="friendlist-image-info">
                                <img src={user.profilePicture} alt={user.username} className="leaderboard-picture"/>
                                <div className="friendlist-info">
                                    <h2 className="friendlist-username">{user.username}</h2>
                                    <p className="friendlist-rating">Rating: {user.rating}</p>
                                </div>
                            </div>
                            <div className={"onlineStatusAndRemoveButton"}>
                                <img src={online ? onlineImage : offlineImage} className="online-status"
                                     alt={"online/offline status"}/>
                                <Button variant={"contained"} color={"error"} onClick={() => removeFromFriendlist(user.username)}>Remove</Button>
                            </div>
                        </li>
                    ))}
                </ul>
            </div>
        </>
    );
}
