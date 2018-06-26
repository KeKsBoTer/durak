import * as React from 'react';

import { Button, Input, Modal } from 'antd';

import 'antd/dist/antd.css';  // or 'antd/dist/antd.less'
import icon from './icon.png';
import './style.css';

const host: string = "http://localhost:8080"

interface ISignUpModal {
  onLogin(user: string): void;
}

class SignUpModal extends React.Component<ISignUpModal, any> {

  private userNameInput: Input | null;

  constructor(props: ISignUpModal) {
    super(props);
    this.state = {
      loading: false,
      user: "",
      visible: false
    }
  }

  public componentDidMount() {
    this.loadUsername()
  }
  public componentWillUnmount() {
    this.handleCancel()
  }

  public render() {
    return (
      <Modal
        visible={this.state.visible}
        onOk={this.handleOk}
        closable={false}
        footer={[
          <Button key="submit" type="primary" loading={this.state.loading} onClick={this.handleOk}>
            Start Playing
            </Button>,
        ]}
      >
        <div className="sign-up-modal">
          <img src={icon} />
          <h1>Wellcome to Durak</h1>
          <h2>Choose your player name:</h2>
        </div>
        <Input size="large" onKeyDown={this.handleInput} ref={(ref) => this.userNameInput = ref} />
      </Modal>
    );
  }

  private handleInput = (e: any) => {
    if (this.userNameInput !== null) {
      this.setState({
        user: this.userNameInput.input.value
      })
    }
  }

  private showModal() {
    this.setState({
      visible: true,
    });
  }

  private handleOk = () => {
    this.setState({ loading: true })
    this.login()
      .then(res => this.setState({ loading: false, visible: false }))
      .then(() => this.props.onLogin(this.state.user))
      .catch(() => this.setState({ loading: false }))
  }
  private handleCancel = () => {
    this.setState({ visible: false });
  }

  private loadUsername = () => {
    fetch(host + "/username", {
      credentials: "include",
      mode: "cors"
    })
      .then(response => {
        switch (response.status) {
          case 200:
            if (response.body != null) {
              response.text()
                .then((text) => {
                  this.setState({ user: text })
                  this.props.onLogin(text)
                })
            } else {
              throw Error("empty body")
            }
            break
          case 401:
            this.showModal()
            return
          default:
            throw Error("invalid status code")
        }
      })
      .catch(((e) => alert(e)))
  }

  private login = () => {
    return fetch(host + "/login?username=" + this.state.user)
  }
}

export default SignUpModal;
