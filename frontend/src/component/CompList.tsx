import React from "react";

export const CompList = (props: any) => {
	const {compTitle, compContent, onClickBack, onClickDelete} = props;
	return (
      <div>
        <ul>
        {compTitle.map((_: string, index: number) => {
          return (
            <li>{compTitle[index]}
              <p>{compContent[index]}</p>
                <button onClick={() => onClickBack(index, false)}>back</button>
                <button onClick={() => onClickDelete(index, false)}>delete</button>
            </li>
          );
        })}
        </ul>
      </div>
	)
};