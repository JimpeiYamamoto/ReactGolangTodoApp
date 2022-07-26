import React from "react";
import { useState } from "react";
import { AddMemo } from "./component/AddMemo"
import { TodoList } from "./component/TodoList"
import { CompList } from "./component/CompList"
import axios from "axios";

const dropElementList = (index: number, list: string[]): string[] => {
  const newList = [...list];
  newList.splice(index, 1);
  return newList;
}

const appendList = (elem: string, list: string[]): string[] => {
  const newList = [...list, elem];
  return newList;
}

export const App = () => {

  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [todoTitle, setTodoTitle] = useState<string[]>([]);
  const [todoContent, setTodoContent] = useState<string[]>([]);
  const [compTitle, setCompTitle] = useState<string[]>([]);
  const [compContent, setCompContent] = useState<string[]>([]);

  const onChangeTitle = (event: any) => setTitle(event.target.value);
  const onChangeContent = (event: any) => setContent(event.target.value);
 
  const onClickAdd = () => {
    if (title===''||content==='') return;
    setTodoTitle(appendList(title, todoTitle))
    setTodoContent(appendList(content, todoContent));
    axios
      .get('http://localhost:8080/addtodo?title=a&content=bbb', {
        headers: {'Content-Type': 'application/json'}, responseType: 'json'
      })
      .then(response => console.log('response body:', response.data));
    setTitle('');
    setContent('');
  }

  const onClickMove = (index: number, isTodoList: boolean) => {
    if (isTodoList) {
      setCompTitle(appendList(todoTitle[index], compTitle));
      setCompContent(appendList(todoContent[index], compContent));
      setTodoTitle(dropElementList(index, todoTitle));
      setTodoContent(dropElementList(index, todoContent));
    } else {
      setTodoTitle(appendList(compTitle[index], todoTitle));
      setTodoContent(appendList(compContent[index], todoContent));
      setCompTitle(dropElementList(index, compTitle));
      setCompContent(dropElementList(index, compContent));
    }
  }

  const onClickDelete = (index: number, isTodoList: boolean) => {
    if (isTodoList) {
      setTodoTitle(dropElementList(index, todoTitle));
      setTodoContent(dropElementList(index, todoContent));
    } else {
      setCompTitle(dropElementList(index, compTitle));
      setCompContent(dropElementList(index, compContent));
    }
  }

  return (
    <>
      <h1>My Memo Todo App</h1>
      <h2>Add Todo</h2>
        <AddMemo 
          todoTitle={title}
          onChangeTitle={onChangeTitle}
          todoContent={content}
          onChangeContent={onChangeContent}
          onClick={onClickAdd}
        />
      <h2>Todo</h2>
        <TodoList 
          todoTitle={todoTitle}
          todoContent={todoContent}
          onClickDone={onClickMove}
          onClickDelete={onClickDelete}
        />
      <h2>Complete</h2>
        <CompList
          compTitle={compTitle}
          compContent={compContent}
          onClickBack={onClickMove}
          onClickDelete={onClickDelete}
        />
    </>
  )
}

