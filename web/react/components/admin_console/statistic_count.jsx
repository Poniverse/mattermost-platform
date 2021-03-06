// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

import {FormattedMessage} from 'mm-intl';

export default class StatisticCount extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        let loading = (
            <FormattedMessage
                id='admin.analytics.loading'
                defaultMessage='Loading...'
            />
        );

        return (
            <div className='col-sm-3'>
                <div className='total-count'>
                    <div className='title'>
                        {this.props.title}
                        <i className={'fa ' + this.props.icon}/>
                    </div>
                    <div className='content'>{this.props.count == null ? loading : this.props.count}</div>
                </div>
            </div>
        );
    }
}

StatisticCount.propTypes = {
    title: React.PropTypes.string.isRequired,
    icon: React.PropTypes.string.isRequired,
    count: React.PropTypes.number
};
