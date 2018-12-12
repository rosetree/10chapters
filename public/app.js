class ChaptersListItem extends React.Component {
    render() {
        return (
            <li className="chapters-list__item">
                <b className="chapters-list__chapter">{this.props.name}</b>&nbsp;
                ({this.props.nr_in_list} / {this.props.chapters_in_list})
            </li>
        );
    }
}

class ChaptersList extends React.Component {
    render() {
        return (
            <ol className="chapters-list">
                {this.props.chapters.map((chapter, index) => (
                    <ChaptersListItem
                        key={chapter.Nr}
                        name={chapter.Chapter}
                        nr_in_list={chapter.NrInList}
                        chapters_in_list={chapter.ChaptersInList} />
                ))}
            </ol>
        );
    }
}

class App extends React.Component {

    state = {
        chapters: null,
        loading: true
    };

    loadChapters = () => {
        // TODO: don’t use hard coded params
        const res = fetch("list?started=2017-12-06&skipped=7&format=json");

        this.setState({ loading: true });

        res
            .then(res => res.json())
            .then(res => {
                this.setState({ chapters: res.Chapters, loading: false });
            });
    };

    componentDidMount() {
        this.loadChapters();
    }

    render() {
        if (this.state.loading) {
            return (
                    <p>Loading.</p>
            );
        }

        return (
                <ChaptersList chapters={this.state.chapters} />
        );
    }
}

const rootElement = document.getElementById("app_root");
ReactDOM.render(<App />, rootElement);
