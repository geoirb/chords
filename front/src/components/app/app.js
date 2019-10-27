import React, { Component } from 'react';
import { Form, Button, Segment} from 'semantic-ui-react';

import GoService from '../../services/go-server';


const { getChords } = new GoService();

class App extends Component{
    state = {
        view: false,
        titel: '',
        author: '',
        text: [],
    }

    onGetText = () => {
        const { titel, author } = this.state;
        getChords(author.replace(/ /g,'+'), titel.replace(/ /g,'+')).then(({ data }) => this.setState({ text: data }));
    };

    onView = (view) => () => this.setState({ view: !view });

    onChange = (e, { name, value }) => this.setState({ [name]: value });
    
    render(){
        const { titel, author, text, view } = this.state;
        
        const splice = (idx, rem, str) => {
            return str.slice(0, idx) + rem + str.slice(idx);
        };

        const getLine = (len, chords) => {
            let chordsLine = ' '.repeat(len)
            chords && chords.forEach(({chord, position}) => chordsLine = splice(position, chord, chordsLine))
            return view ? <b>{chordsLine}</b>: null
        }

        return (
            <Segment>
                <Form>
                    <Form.Group widths='equal'>
                        <Form.Input
                            onChange={this.onChange}
                            value={titel}
                            name='titel'
                            label='Название песни'
                            placeholder='Введите название песни'
                        />
                        <Form.Input
                            onChange={this.onChange}
                            value={author}
                            name='author'
                            label='Артист'
                            placeholder='Введите артиста'
                        />
                    </Form.Group>
                </Form>
                <Button
                    onClick={this.onGetText}
                    icon='check'
                    disabled={author === "" || titel === ""}
                    />
                <Button 
                    positive={view}
                    negative={!view}
                    onClick={this.onView(view)}
                    icon='eye'
                    />
                <Segment>
                    <pre>
                    {
                        text.map(({line, chords})=>(
                            <React.Fragment> 
                                {getLine(line.length, chords)}
                                {line !== "" ? <div>{line}</div> : <p></p>}
                            </React.Fragment>
                        ))
                    }
                    </pre>
                </Segment>
            </Segment>
        );
    };
};

export default App;

