import React, { useState } from "react";
import SimpleBox from "./SimpleBox";

export default function Column({ onClick, columnKey, diceRoll }: { onClick: Function, columnKey: number, diceRoll: number | null }) {
    const [boxes, setBoxes] = useState<(number | null)[]>([null, null, null]);

    const handleClick = () => {
        console.log(`Column ${columnKey} clicked`);
        if (onClick && diceRoll !== null) {
            onClick(columnKey);
            const newBoxes = [...boxes];
            for (let i = newBoxes.length - 1; i >= 0; i--) {
                if (newBoxes[i] === null) {
                    newBoxes[i] = diceRoll;
                    break;
                }
            }
            setBoxes(newBoxes);
        }
    };

    return (
        <div onClick={handleClick} style={{ cursor: "pointer" }}>
            {boxes.map((box, index) => (
                <SimpleBox key={index} diceValue={box} />
            ))}
        </div>
    );
}
