import React, { createContext, useContext, useState, ReactNode } from 'react';
import ReactAudioPlayer from 'react-audio-player';
import backgroundMusic from '../soundtracks/backgroundMusic.mp3';

interface AudioContextProps {
    isPlaying: boolean;
    setIsPlaying: React.Dispatch<React.SetStateAction<boolean>>;
}

const AudioContext = createContext<AudioContextProps | undefined>(undefined);

interface AudioProviderProps {
    children: ReactNode;
}

export const AudioProvider: React.FC<AudioProviderProps> = ({ children }) => {
    const [isPlaying, setIsPlaying] = useState(true);

    return (
        <AudioContext.Provider value={{ isPlaying, setIsPlaying }}>
            <ReactAudioPlayer
                src={backgroundMusic}
                autoPlay={true}
                loop={true}
                preload="auto"
                controls
            />
            {children}
        </AudioContext.Provider>
    );
};

export const useAudio = (): AudioContextProps => {
    const context = useContext(AudioContext);
    if (!context) {
        throw new Error('useAudio must be used within an AudioProvider');
    }
    return context;
};