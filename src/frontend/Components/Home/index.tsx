import * as React from 'react';
import cx from 'classnames';
import subDays from 'date-fns/subDays';
import isBefore from 'date-fns/isBefore';
import Spinner from 'react-spinkit';

import { ApiClient, NotebookHandle, Recipe } from "../../types";

interface NotebooksListProps {
  notebooks: NotebookHandle[];
}

function NotebooksList(props: NotebooksListProps) {
  const { notebooks } = props;
  return (
    <div className="notebook-list">
      <ul>
        {notebooks.map(notebook => (
          <li key={notebook.name}>
            <a href={notebook.url}>{notebook.name}</a>
            <span className="notebook-recipe">{notebook.recipe.name}</span>
          </li>
        ))}
      </ul>
    </div>
  );
}

export interface Props {
  apiClient: ApiClient;
  notebooks: NotebookHandle[];
  recipes: Recipe[];
  newnotebookurl: string;
}

interface State {
  menuopen: boolean;
  creating: boolean;
}

export default class Home extends React.Component<Props, State> {

  state: State = {
    menuopen: false,
    creating: false,
  };

  declare props: Props;

  render() {

    const { notebooks, recipes } = this.props;
    const { menuopen, creating } = this.state;

    const recents: NotebookHandle[] = [];
    const horizon = subDays(new Date(), 1);

    notebooks.forEach(notebook => {
      if (!isBefore(new Date(notebook.mtime), horizon)) {
        recents.push(notebook);
      }
    });

    recents.sort((a, b) => a.mtime > b.mtime ? -1 : 1);

    const showTitle = recents.length > 0;

    return (
      <div className="home-app">

        <div className="list">
          {recents.length > 0 && (
            <div>
              <h2>最近({recents.length})</h2>
              <NotebooksList notebooks={recents} />
            </div>
          )}

          {notebooks.length > 0 && notebooks.length > recents.length && (
            <div>
              {showTitle && <h2>全部({notebooks.length})</h2>}
              <NotebooksList notebooks={notebooks} />
            </div>
          )}
        </div>

        <div className="tools">
          <button className={cx('bigbutton', 'btn-new', { creating })}
            onClick={() => (this as any).setState({ menuopen: !menuopen })}>
            {creating ? <Spinner fadeIn="none" name="wave" color="white" /> : (
              <span>{menuopen ? '-' : '+'} 笔记本</span>
            )}
          </button>
          {menuopen && (
            <div className="recipe-list">{recipes.map(recipe => (
              <span key={recipe.key} className="recipe-item"
                onClick={() => this.selectRecipe(recipe.key)}>{recipe.name}</span>
            ))}</div>
          )}
        </div>
      </div>
    );
  }

  private selectRecipe(recipekey: string) {
    const { newnotebookurl } = this.props;
    const { creating } = this.state;
    if (creating) return;

    (this as any).setState({ creating: true });

    this.props.apiClient.create(newnotebookurl, recipekey)
      .then(res => res.json())
      .then(({ url }) => {
        document.location.href = url;
      })
      .catch(_ => {
        alert('Error: Notebook could not be created.');
      });
  }
}
