import React, { Component } from 'react';
import actions from '../actions/urls.actions';
import { connect } from 'react-redux';
import Container from 'react-bootstrap/Container';
import Form from 'react-bootstrap/Form';
import _ from 'lodash';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Button from 'react-bootstrap/Button';
import InputGroup from 'react-bootstrap/InputGroup';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlus, faTrash } from '@fortawesome/free-solid-svg-icons';
import classNames from 'classnames';

class Urls extends Component {
  constructor(props) {
    super(props);

    this.state = { urls: {} };
  }

  componentDidMount() {
    this.props.dispatch(actions.getUrls);
  }

  handleChange(short, e) {
    const urlLong = e.target.value;
    this.setState(prevState => {
      prevState.urls[short] = urlLong;
      return prevState;
    });
  }

  submitChange(short, e) {
    this.props.dispatch(actions.setUrl(short, e.target.value));
  }

  handleAddUrl(e) {
    const elements = _.get(e, 'currentTarget.elements');
    if (!_.isNil(elements)) {
      const short = elements[0].value;
      const long = elements[1].value;
      this.props.dispatch(actions.setUrl(short, long));
      elements[0].value = '';
      elements[1].value = '';
    }
  }

  componentDidUpdate(prevProps, prevState, snapshot) {
    if (!_.isEqual(prevProps, this.props)) {
      this.setState(() => {
        return {
          urls: _.transform(
            this.props.urls,
            (result, url) => {
              result[url.short] = _.cloneDeep(url.long);
            },
            {}
          )
        };
      });
    }
  }

  render() {
    return (
      <Container
        className={classNames('d-flex', 'flex-column', 'align-content-center')}
        fluid
      >
        <h1>Url Links</h1>
        <h2>Admin tool</h2>
        {_.map(this.state.urls, (url, short) => {
          return (
            <Row
              style={{ 'margin-bottom': 5 }}
              className={classNames('d-flex', 'align-content-center')}
            >
              <Form key={`${short}-form`} id={`${short}-form`} inline>
                <Col md={'auto'}>
                  <Form.Control
                    key={`${short}-label-control`}
                    value={short}
                    disabled
                  />
                </Col>

                <Col md={'auto'}>
                  <InputGroup>
                    <Form.Control
                      key={`${short}-input`}
                      id={`${short}-input`}
                      as={'input'}
                      value={this.state.urls[short]}
                      onChange={e => {
                        this.handleChange(short, e);
                      }}
                      onBlur={e => {
                        this.submitChange(short, e);
                      }}
                      type={'text'}
                    />
                    <InputGroup.Append>
                      <Button
                        variant={'danger'}
                        onClick={() => {
                          this.props.dispatch(
                            actions.removeUrl(short, this.state.urls[short])
                          );
                        }}
                      >
                        <FontAwesomeIcon icon={faTrash} />
                      </Button>
                    </InputGroup.Append>
                  </InputGroup>
                </Col>
              </Form>
            </Row>
          );
        })}
        <Row>
          <label>Add Url Mapping:</label>
        </Row>
        <Row>
          <Form
            onSubmit={e => {
              e.preventDefault();
              this.handleAddUrl(e);
            }}
            inline
          >
            <Col md={'auto'}>
              <Form.Control type={'text'} placeholder={'Enter Url key'} />
            </Col>

            <Col md={'auto'}>
              <InputGroup>
                <Form.Control type={'text'} placeholder={'Enter Url value'} />
                <InputGroup.Append>
                  <Button variant={'primary'} type='submit'>
                    <FontAwesomeIcon icon={faPlus} />
                  </Button>
                </InputGroup.Append>
              </InputGroup>
            </Col>
          </Form>
        </Row>
      </Container>
    );
  }
}

const mapStateToProps = state => {
  return { urls: state.urls };
};

export default connect(mapStateToProps)(Urls);
