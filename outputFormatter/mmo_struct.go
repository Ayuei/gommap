package outputFormatter

import (
	"encoding/xml"
	"time"
)

/*
Struct that matches the structure of MetaMap's XML output.
Output from: https://www.onlinetool.io/xmltogo/
*/

type MMOs struct {
	XMLName xml.Name `xml:"MMOs"`
	Text    string   `xml:",chardata"`
	MMO     struct {
		Text    string `xml:",chardata"`
		CmdLine struct {
			Text    string `xml:",chardata"`
			Command string `xml:"Command"`
			Options struct {
				Text   string `xml:",chardata"`
				Count  string `xml:"Count,attr"`
				Option []struct {
					Text     string `xml:",chardata"`
					OptName  string `xml:"OptName"`
					OptValue string `xml:"OptValue"`
				} `xml:"Option"`
			} `xml:"Options"`
		} `xml:"CmdLine"`
		AAs struct {
			Text  string `xml:",chardata"`
			Count string `xml:"Count,attr"`
		} `xml:"AAs"`
		Negations struct {
			Text  string `xml:",chardata"`
			Count string `xml:"Count,attr"`
		} `xml:"Negations"`
		Utterances struct {
			Text      string `xml:",chardata"`
			Count     string `xml:"Count,attr"`
			Utterance []struct {
				Text        string `xml:",chardata"`
				PMID        string `xml:"PMID"`
				UttSection  string `xml:"UttSection"`
				UttNum      string `xml:"UttNum"`
				UttText     string `xml:"UttText"`
				UttStartPos string `xml:"UttStartPos"`
				UttLength   string `xml:"UttLength"`
				Phrases     struct {
					Text   string `xml:",chardata"`
					Count  string `xml:"Count,attr"`
					Phrase []struct {
						Text        string `xml:",chardata"`
						PhraseText  string `xml:"PhraseText"`
						SyntaxUnits struct {
							Text       string `xml:",chardata"`
							Count      string `xml:"Count,attr"`
							SyntaxUnit []struct {
								Text       string `xml:",chardata"`
								SyntaxType string `xml:"SyntaxType"`
								LexMatch   string `xml:"LexMatch"`
								InputMatch string `xml:"InputMatch"`
								LexCat     string `xml:"LexCat"`
								Tokens     struct {
									Text  string   `xml:",chardata"`
									Count string   `xml:"Count,attr"`
									Token []string `xml:"Token"`
								} `xml:"Tokens"`
							} `xml:"SyntaxUnit"`
						} `xml:"SyntaxUnits"`
						PhraseStartPos int `xml:"PhraseStartPos"`
						PhraseLength   int `xml:"PhraseLength"`
						Candidates     struct {
							Text      string `xml:",chardata"`
							Total     int `xml:"Total,attr"`
							Excluded  int `xml:"Excluded,attr"`
							Pruned    int `xml:"Pruned,attr"`
							Remaining int `xml:"Remaining,attr"`
						} `xml:"Candidates"`
						Mappings struct {
							Text    string `xml:",chardata"`
							Count   string `xml:"Count,attr"`
							Mapping []struct {
								Text              string `xml:",chardata"`
								MappingScore      string `xml:"MappingScore"`
								MappingCandidates struct {
									Text      string `xml:",chardata"`
									Total     string `xml:"Total,attr"`
									Candidate []struct {
										Text               string `xml:",chardata"`
										CandidateScore     string `xml:"CandidateScore"`
										CandidateCUI       string `xml:"CandidateCUI"`
										CandidateMatched   string `xml:"CandidateMatched"`
										CandidatePreferred string `xml:"CandidatePreferred"`
										MatchedWords       struct {
											Text        string   `xml:",chardata"`
											Count       string   `xml:"Count,attr"`
											MatchedWord []string `xml:"MatchedWord"`
										} `xml:"MatchedWords"`
										SemTypes struct {
											Text    string `xml:",chardata"`
											Count   string `xml:"Count,attr"`
											SemType []string `xml:"SemType"`
										} `xml:"SemTypes"`
										MatchMaps struct {
											Text     string `xml:",chardata"`
											Count    string `xml:"Count,attr"`
											MatchMap []struct {
												Text           string `xml:",chardata"`
												TextMatchStart string `xml:"TextMatchStart"`
												TextMatchEnd   string `xml:"TextMatchEnd"`
												ConcMatchStart string `xml:"ConcMatchStart"`
												ConcMatchEnd   string `xml:"ConcMatchEnd"`
												LexVariation   string `xml:"LexVariation"`
											} `xml:"MatchMap"`
										} `xml:"MatchMaps"`
										IsHead      string `xml:"IsHead"`
										IsOverMatch string `xml:"IsOverMatch"`
										Sources     struct {
											Text   string   `xml:",chardata"`
											Count  string   `xml:"Count,attr"`
											Source []string `xml:"Source"`
										} `xml:"Sources"`
										ConceptPIs struct {
											Text      string `xml:",chardata"`
											Count     string `xml:"Count,attr"`
											ConceptPI []struct {
												Text     string `xml:",chardata"`
												StartPos string `xml:"StartPos"`
												Length   string `xml:"Length"`
											} `xml:"ConceptPI"`
										} `xml:"ConceptPIs"`
										Status  string `xml:"Status"`
										Negated string `xml:"Negated"`
									} `xml:"Candidate"`
								} `xml:"MappingCandidates"`
							} `xml:"Mapping"`
						} `xml:"Mappings"`
					} `xml:"Phrase"`
				} `xml:"Phrases"`
			} `xml:"Utterance"`
		} `xml:"Utterances"`
	} `xml:"MMO"`
	ParseTime time.Duration
	RawXML    string
	ItemID string
}
