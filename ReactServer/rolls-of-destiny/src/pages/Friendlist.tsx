import React, {useEffect, useState} from 'react';
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
import {authFetch} from "../auth";

export default function Friendlist({loggedIn}: { loggedIn: boolean }) {
    const [searchBar, setSearchBar] = useState("");
    const [searchResults, setSearchResults] = useState<profile[]>([]);
    const [resultsFound, setResultsFound] = useState(false);
    const [noResultsFound, setNoResultsFound] = useState(false);

    const [users, setUsers] = useState<profile[]>([
    ])

    const online = true;


    function submitSearchBar() {
        // for connecting with backend
        authFetch("http://menews.site:9090/getAccounts?username=" + searchBar).then(r => {
            if (r.status === 200) {
                r.json().then(r => {
                    console.log(r)
                    setSearchResults(r.friends);
                    if (r.friends.length > 0) {
                        setResultsFound(true);
                        setNoResultsFound(false);
                    } else {
                        setResultsFound(false);
                        setNoResultsFound(true);
                    }
                })
            }
        });
    }

    const addToFriendlist = (username: string) => {
        authFetch("http://menews.site:9090/addFriend", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({username: username})
        })
    };

    useEffect(() => {
            console.log("test")
            authFetch("http://menews.site:9090/getFriends").then(response => {
                console.log(response); return response.json()
            }).then(response => {
                console.log(response)
                setUsers(response.friends)
            })
        }
        ,
        []
    )
    const handleCloseSearchModul = () => {
        setSearchBar("");
        setSearchResults([]);
        setResultsFound(false);
        setNoResultsFound(false)
        window.location.reload();
    };
    const removeFromFriendlist = (username: string) => {
        authFetch("http://menews.site:9090/removeFriend", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({username: username})
        }).then(r => {
            window.location.reload();
        })
    };

    return (
        <>
            <TopAppBar loggedIn={loggedIn}/>
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
                                <img src={"data:image/jpeg;base64," + user.profilePicture} alt={user.username} className="leaderboard-picture"/>
                                <div className="friendlist-info">
                                    <h2 className="friendlist-username">{user.username}</h2>
                                    <p className="friendlist-rating">Rating: {user.rating}</p>
                                </div>
                            </div>
                            <div className={"onlineStatusAndRemoveButton"}>
                                <img src={online ? onlineImage : offlineImage} className="online-status"
                                     alt={"online/offline status"}/>
                                <Button variant={"contained"} color={"error"}
                                        onClick={() => removeFromFriendlist(user.username)}>Remove</Button>
                            </div>
                        </li>
                    ))}
                </ul>
            </div>
        </>
    );
}
