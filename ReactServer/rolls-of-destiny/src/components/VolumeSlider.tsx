import {VolumeDown, VolumeUp} from "@mui/icons-material";
import {Slider, Stack} from "@mui/material";
import React, {useState} from "react";


export default function VolumeSlider() {

    const initialVolume = Number(sessionStorage.getItem("volume")) || 99;
    const [volume, setVolume] = useState<number>(initialVolume);

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