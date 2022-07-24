import React from "react";

export const TodoList = (props: any) => {
	const { todoTitle, todoContent, onClickDone, onClickDelete }  = props;
	return (
      <div>
        <ul>
          {todoTitle.map((_: string, index:number) => {
            return (
              <li>{todoTitle[index]}
                <p>{todoContent[index]}</p>
                <button onClick={() => onClickDone(index, true)}>done</button>
                <button onClick={() => onClickDelete(index, true)}>delete</button>
              </li>
            );
          })}
        </ul>
      </div>
	)
};