package usecase

type Container struct {
	saveDefinition        SaveDefinitionUseCase
	getLessonHomework     GetLessonHomeworkUseCase
	getStudentHomework    GetStudentHomeworkUseCase
	submitHomework        SubmitHomeworkUseCase
	listReviewSubmissions ListReviewSubmissionsUseCase
	getReviewSubmission   GetReviewSubmissionUseCase
	reviewSubmission      ReviewSubmissionUseCase
}

func NewContainer(saveDefinition SaveDefinitionUseCase, getLessonHomework GetLessonHomeworkUseCase, getStudentHomework GetStudentHomeworkUseCase, submitHomework SubmitHomeworkUseCase, listReviewSubmissions ListReviewSubmissionsUseCase, getReviewSubmission GetReviewSubmissionUseCase, reviewSubmission ReviewSubmissionUseCase) *Container {
	return &Container{saveDefinition: saveDefinition, getLessonHomework: getLessonHomework, getStudentHomework: getStudentHomework, submitHomework: submitHomework, listReviewSubmissions: listReviewSubmissions, getReviewSubmission: getReviewSubmission, reviewSubmission: reviewSubmission}
}

func (c *Container) SaveDefinition() SaveDefinitionUseCase { return c.saveDefinition }

func (c *Container) GetLessonHomework() GetLessonHomeworkUseCase { return c.getLessonHomework }

func (c *Container) GetStudentHomework() GetStudentHomeworkUseCase { return c.getStudentHomework }

func (c *Container) SubmitHomework() SubmitHomeworkUseCase { return c.submitHomework }

func (c *Container) ListReviewSubmissions() ListReviewSubmissionsUseCase {
	return c.listReviewSubmissions
}

func (c *Container) GetReviewSubmission() GetReviewSubmissionUseCase { return c.getReviewSubmission }

func (c *Container) ReviewSubmission() ReviewSubmissionUseCase { return c.reviewSubmission }
