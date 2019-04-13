class SimpleAxiosGet extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            open: false,
            connected: false,
        };
        this.wspath = location.protocol.match(/^https/) ? "wss://" : "ws://";
        this.socket = new WebSocket(this.wspath + location.hostname + ':' + location.port + '/ws');
        this.socket.onopen = () => {
            this.setState({connected: true})
        };

    }

    componentDidMount() {
        this.socket.onmessage = ({ data }) => {
            let jdata = JSON.parse(data);
            console.log(jdata);
        }
    }

    render() {
        let who = "World";
        return (
            <div>
                <div>
                    <h1> Hello {who}! </h1>
                </div>
            </div>
        );
    }
}

ReactDOM.render( <SimpleAxiosGet/>, document.querySelector("#root"));