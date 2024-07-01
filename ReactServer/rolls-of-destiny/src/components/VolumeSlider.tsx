import {VolumeDown, VolumeUp} from "@mui/icons-material";
import {Slider, Stack} from "@mui/material";
import React, {useState} from "react";


export default function VolumeSlider({volume, setVolume} : {volume: number, setVolume: Function}) {


    const handleVolumeChange = (event: Event, newValue: number | number[]) => {
        setVolume(newValue as number);
        sessionStorage.setItem("volume", String(volume));
    };

    return (
        <Stack spacing={2} direction="row" sx={{ mb: 1 }} alignItems="center">
            <VolumeDown />
            <Slider value={volume} aria-label="Volume" onChange={handleVolumeChange} />
            <VolumeUp />
        </Stack>
    )
}