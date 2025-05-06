package competition

import "yadro-test-assigment/internal"

// Ищет или создает нового участника
func (receiver *Competition) findCompetitor(CompetitorID int) *internal.Competitor {
	if receiver.Competitors[CompetitorID] == nil {
		receiver.Competitors[CompetitorID] = &internal.Competitor{
			ID: CompetitorID,
		}
	}
	return receiver.Competitors[CompetitorID]
}
