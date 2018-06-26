import * as React from 'react';

import { Button, Card } from 'antd'

import 'antd/dist/antd.css';  // or 'antd/dist/antd.less'

import styles from './JoinGame.css'

class JoinGame extends React.Component<any, any> {

    constructor(props: any) {
        super(props);
    }

    public render() {
        return (
            <Card className={styles.container}>
                <h1>Press button to search for a game.</h1>
                <Button>Play</Button>
            </Card>
        );
    }
}

export default JoinGame;
