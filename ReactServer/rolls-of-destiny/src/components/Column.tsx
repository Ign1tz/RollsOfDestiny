import React, {useState} from "react";
import SimpleBox from "./SimpleBox";

export default function Column({canPlace, setCanPlace, columnKey, diceRoll}: {
    canPlace?: boolean,
    setCanPlace?: Function,
    columnKey: number,
    diceRoll: number | null
}) {
    const [boxes, setBoxes] = useState<(number | null)[]>([null, null, null]);

    const handleClick = () => {
        if (canPlace && diceRoll !== null && setCanPlace) {
            console.log("DiceRoll received and CanPlace True")
            const newBoxes = [...boxes];
            for (let i = newBoxes.length - 1; i >= 0; i--) {
                if (newBoxes[i] === null) {
                    newBoxes[i] = diceRoll;
                    break;
                }
            }
            setBoxes(newBoxes);
            setCanPlace(false);
        }
    };

    return (
        <div onClick={handleClick} style={{cursor: "pointer"}}>
            {boxes.map((box, index) => (
                <SimpleBox key={index} diceValue={box}/>
            ))}
        </div>
    );
}
