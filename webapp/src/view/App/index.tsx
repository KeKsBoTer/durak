import * as React from 'react';

import JoinGame from '../../components/JoinGame';
import SignUpModal from '../../components/SignUpModal';

import 'antd/dist/antd.css';  // or 'antd/dist/antd.less'
import './style.css';

interface IAppState {
  user?: string;
}

class App extends React.Component<any, IAppState> {

  constructor(props: any) {
    super(props);
    this.state = {
      user: undefined
    }
  }

  public render() {
    return (
      <div>
        {this.state.user &&
          <JoinGame/>
        }
        <SignUpModal onLogin={this.onLogin} />
      </div>
    );
  }

  private onLogin = (user: string) => {
    this.setState({user})
  }
}

export default App;
