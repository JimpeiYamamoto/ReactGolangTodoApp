import React from "react";

export const AddMemo = (props: any) => {
	const { todoTitle, onChangeTitle, todoContent, onChangeContent, onClick} = props;
	return (
      <div>
        <input placeholder="title" value={todoTitle} onChange={onChangeTitle}/>
        <input placeholder="content" value={todoContent} onChange={onChangeContent}/>
        <button onClick={onClick}>Add</button>
      </div>
	)
};